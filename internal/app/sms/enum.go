package sms

type SmsStatus int

const (
	SmsStatusSent SmsStatus = iota
	SmsStatusFailed
)

func (e SmsStatus) String() string {
	return [...]string{"sent", "failed"}[e]
}
