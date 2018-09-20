package channel

const (
	Account_source = "账号提醒"
	Sms_source     = "短信提醒"
	Email_source   = "邮件提醒"
	App_source     = "APP提醒"
)

const (
	No_enouth_10    = "不足10次"
	No_enouth_50    = "不足50次"
	No_enouth_day_3 = "三天内到期"
	No_enouth_day_7 = "七天内到期"
	No_enouth_day_1 = "当天到期"
)

type ALertMeta struct {
	AleFrom string `json:"alert_source"`
	AleType string `json:"alert_type"`
	AleTo   string `json:"alert_to"`
}

func (a *ALertMeta) Euqal(b *ALertMeta) bool {
	if a.AleFrom == b.AleFrom && a.AleType == b.AleType && a.AleTo == b.AleTo {
		return true
	}
	return false
}
