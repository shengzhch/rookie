package channel

import (
	"de-tcp/util"
	"fmt"
)

var (
	channelFactories         = make(map[string]IChannelFactory)
	channelObserverFactories = make(map[string]IChannelObserverFactory)
	ruleFactories            = make(map[string]IRuleFactory)
)

func RegisChannelFactory(name string, cf IChannelFactory) {
	channelFactories[name] = cf
}

func RegisChannelObserverFactory(name string, cof IChannelObserverFactory) {
	channelObserverFactories[name] = cof
}

func RegisRuleFactory(name string, rf IRuleFactory) {
	ruleFactories[name] = rf
}

func NewChannel(cof *util.Config) (IChannel, error) {
	if cof == nil {
		return nil, NewConfigErr(nil, "config is nil")
	}
	cf_name := cof.Get("CF_NAME").MustString()
	fmt.Println("CF_NAME", cf_name)
	return channelFactories[cf_name].New(cof)
}

func NewChannelObserver(cof *util.Config) (IChannelObserver, error) {
	if cof == nil {
		return nil, NewConfigErr(nil, "config is nil")
	}
	cof_name := cof.Get("COF_NAME").MustString()
	fmt.Println("COF_NAME", cof_name)
	return channelObserverFactories[cof_name].New(cof)
}

func NewRule(cof *util.Config) (IRule, error) {
	if cof == nil {
		return nil, NewConfigErr(nil, "config is nil")
	}
	rf_name := cof.Get("RF_NAME").MustString()
	fmt.Println("RF_NAME", rf_name)
	return ruleFactories[rf_name].New(cof)
}
