package messagebird

import (
	"airbnb-messaging-be/internal/pkg/env"

	"github.com/messagebird/go-rest-api/v9/sms"
)

func SendSms(recipients []string, body string, msgParams *sms.Params) (*sms.Message, error) {
	return sms.Create(MessengerClient, env.CONFIG.Messenger.Originator, recipients, body, msgParams)
}
