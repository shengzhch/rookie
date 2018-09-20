package receive

import (
	"de-tcp/channel"
	"de-tcp/util"
	"fmt"
	"sync"
	"time"
)

type ReceChannel struct {
	*sync.Mutex
	c         *util.Config
	name      string
	over      bool
	firstTime time.Time
	lastTime  time.Time
	timer     *time.Timer
	count     int
	obs       map[string]channel.IChannelObserver
	duration  time.Duration
	lifespan  time.Duration
	rule      channel.IRule
}

type ReceChanneleFactory struct{}

func init() {
	channel.RegisChannelFactory("receive-channel-factory", &ReceChanneleFactory{})
}

func (rcf *ReceChanneleFactory) New(config *util.Config) (channel.IChannel, error) {
	rel := &ReceChannel{
		Mutex:     new(sync.Mutex),
		c:         config,
		name:      "receive",
		over:      false,
		firstTime: time.Now(),
		lastTime:  time.Now(),
		count:     1,
		lifespan:  time.Hour * 48,
		duration:  time.Hour * 24,
		obs:       make(map[string]channel.IChannelObserver),
	}

	rel.run()
	return rel, nil
}

func (rc *ReceChannel) run() {
	f := func() {
		if time.Now().Sub(rc.lastTime) > rc.lifespan {
			rc.over = true
		}
	}
	rc.timer = time.AfterFunc(rc.duration, f)
}

func (rc *ReceChannel) Name() string {
	return rc.name
}

func (rc *ReceChannel) RegisterObserver(co channel.IChannelObserver) {
	rc.Lock()
	rc.obs[co.Name()] = co
	rc.Unlock()
}

func (rc *ReceChannel) UnRegisterObserver(name string) {
	rc.Lock()
	if _, ok := rc.obs[name]; ok {
		delete(rc.obs, name)
	}
	rc.Unlock()
}

func (rc *ReceChannel) ChannelObserver(name string) channel.IChannelObserver {
	return rc.obs[name]
}

func (rc *ReceChannel) StartBySet(config util.Config) {
}

func (rc *ReceChannel) StopBySet() error {
	return nil
}

func (rc *ReceChannel) Can(am *channel.ALertMeta) bool {
	if rc.rule == nil {
		fmt.Println("Error: channel rule is not set")
		return true
	}
	rel := rc.rule.Check(rc.firstTime, rc.lastTime, rc.count)
	rc.Lock()
	if rel {
		rc.lastTime = time.Now()
		rc.count = rc.count + 1
	}
	rc.Unlock()
	return rel
}

func (rc *ReceChannel) HasOver() bool {
	if rc.over {
		rc.timer.Stop()
	}
	return rc.over
}

func (rc *ReceChannel) SetRule(rule channel.IRule) {
	rc.rule = rule
}
