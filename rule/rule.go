package rule

import (
	"de-tcp/channel"
	"de-tcp/util"
	"time"
)

type rule1factort struct {
}

func (r1f *rule1factort) New(config *util.Config) (channel.IRule, error) {
	max := config.Get("max_count").MustInt()
	duration := time.Duration(config.Get("duration").MustInt64())
	if max == 0 || duration == 0 {
		return nil, channel.NewConfigErr([]string{"max_count", "duration"}, "")
	}
	return &Rule1{
		max: max, duration: duration,
	}, nil
}

//多久发送发送一次，最多发送几次
type Rule1 struct {
	max      int
	duration time.Duration
}

func (r1 *Rule1) Check(firsttime, lastime time.Time, count int) bool {
	if time.Now().Sub(lastime) > r1.duration || (count > r1.max && r1.max != 0) {
		return false
	}
	return true
}

func init() {
	channel.RegisRuleFactory("rule_1_factory", &rule1factort{})
}
