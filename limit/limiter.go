package limit

import (
	"de-tcp/channel"
	"de-tcp/util"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type Limiter interface {
	Can(raw []byte) bool
}

type limiter struct {
	*sync.Mutex
	ams      map[string]*channel.ALertMeta
	channels map[string]channel.IChannel
}

func NewLimiter() Limiter {
	l := &limiter{
		Mutex:    new(sync.Mutex),
		ams:      make(map[string]*channel.ALertMeta),
		channels: make(map[string]channel.IChannel),
	}
	go l.run()
	return l
}

func (l *limiter) run() {
	for {
		select {
		case <-time.After(time.Hour):
			l.check()
		}
	}
}

func (l *limiter) check() {
	l.Lock()
	for k, v := range l.channels {
		if v.HasOver() {
			delete(l.ams, k)
			delete(l.channels, k)
		}
	}
	l.Unlock()
}

func (l *limiter) Can(raw []byte) bool {
	am := channel.ALertMeta{}
	err := json.Unmarshal(raw, &am)
	if err != nil {
		fmt.Printf("decode raw byte to alertmata failed %v", err)
		return false
	}
	id := l.GetUUid(&am)
	fmt.Print("id-", id)
	fmt.Printf("%v\n", am)
	return l.can(id, &am)
}

func (l *limiter) can(id string, meta *channel.ALertMeta) bool {
	if c, ok := l.channels[id]; ok {
		return c.Can(meta)
	}
	c := getAconf(meta)
	ch, err := channel.NewChannel(c)
	if err != nil {
		fmt.Printf("new channel failed during check limit is can %v", err)
		return true
	}
	rule, err := channel.NewRule(c)
	if err != nil {
		fmt.Printf("new rule failed during check limit is can %v", err)
		return true
	}
	ch.SetRule(rule)
	go func() {
		l.Lock()
		l.channels[id] = ch
		l.Unlock()
	}()
	return ch.Can(meta)

}

func (l *limiter) GetUUid(meta *channel.ALertMeta) string {
	for k, v := range l.ams {
		if v.Euqal(meta) {
			return k
		}
	}
	id := getAId()
	go func(id string) {
		l.Lock()
		l.ams[id] = meta
		l.Unlock()
	}(id)
	return id

}

func getAId() string {
	return uuid.NewV4().String()
}

//todo 按照提醒类型从数据库中取得规则,生成配置
func getAconf(am *channel.ALertMeta) *util.Config {
	c := util.NewConfig()
	c.Set("CF_NAME", "receive-channel-factory")
	if am.AleFrom == channel.App_source {
		c.Set("RF_NAME", "rule_1_factory")
		c.Set("max_count", 5)
		c.Set("duration", int64(time.Hour*24))
	}
	return c
}
