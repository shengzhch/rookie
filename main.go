package main

import (
	"de-tcp/channel"
	"de-tcp/limit"
	_ "de-tcp/receive"
	_ "de-tcp/rule"
	"encoding/json"
	"fmt"
)

func main() {
	var rel bool
	l := limit.NewLimiter()
	b, _ := json.Marshal(map[string]string{
		"alert_source": channel.App_source,
		"alert_type":   "七天内到期",
		"alert_to":     "123",
	})
	rel = l.Can(b)
	fmt.Println("rel", rel)
	rel = l.Can(b)
	fmt.Println("rel", rel)
	rel = l.Can(b)
	fmt.Println("rel", rel)
	rel = l.Can(b)
	fmt.Println("rel", rel)
	rel = l.Can(b)
	fmt.Println("rel", rel)
	rel = l.Can(b)
	fmt.Println("rel", rel)
	b, _ = json.Marshal(map[string]string{
		"alert_source": channel.App_source,
		"alert_type":   "七天内到期",
		"alert_to":     "1234",
	})
	rel = l.Can(b)
	fmt.Println("rel", rel)

}
