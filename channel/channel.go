package channel

import (
	"de-tcp/util"
	"time"
)

type IChannel interface {
	Name() string
	RegisterObserver(IChannelObserver)
	UnRegisterObserver(string)
	ChannelObserver(string) IChannelObserver
	StartBySet(config util.Config)
	StopBySet() error

	Can(am *ALertMeta) bool
	HasOver() bool
	SetRule(rule IRule)
}

type IChannelObserver interface {
	Name() string
	DataAvailable()
}

//根据第一次请求的时间，上次请求的时间以及这是第几次请求 来确定本次请求是否允许
type IRule interface {
	Check(firsttime time.Time, lastime time.Time, count int) bool
}

type IChannelFactory interface {
	New(*util.Config) (IChannel, error)
}

type IChannelObserverFactory interface {
	New(*util.Config) (IChannelObserver, error)
}

type IRuleFactory interface {
	New(*util.Config) (IRule, error)
}
