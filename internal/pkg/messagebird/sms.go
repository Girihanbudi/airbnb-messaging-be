package messagebird

import (
	"github.com/messagebird/go-rest-api/v9/sms"
)

type SmsParams struct {
	Originator *string
	sms.Params
}

func (m *Messenger) SendSms(recipients []string, body string, originator *string, params *sms.Params) (*sms.Message, error) {
	ori := m.Originator
	if originator != nil {
		ori = *originator
	}
	return sms.Create(m.Client, ori, recipients, body, params)
}
