package channel

import (
	"fmt"
)

type ConfigErr struct {
	MissKeys []string
	Msg      string
}

func (e *ConfigErr) Error() string {
	if e.MissKeys != nil {
		return fmt.Sprintf("Miss necessary key %s\n", e.MissKeys)
	}
	return e.Msg
}

func NewConfigErr(keys []string, msg string) error {
	return &ConfigErr{keys, msg}
}
