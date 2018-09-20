package observer

import (
	"de-tcp/channel"
	"de-tcp/util"
)

type MonitorFactory struct {
}

func (mf *MonitorFactory) New(config *util.Config) (channel.IChannelObserver, error) {
	return &Monitor{}, nil
}

func init() {
	channel.RegisChannelObserverFactory("monitor-observer-factorty", &MonitorFactory{})
}

type Monitor struct {
	name string
}

func (m *Monitor) Name() string {
	return m.name
}

func (m *Monitor) DataAvailable() {
}
