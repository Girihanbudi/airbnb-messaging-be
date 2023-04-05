package sms

import "github.com/messagebird/go-rest-api/v9/sms"

type Payload struct {
	Recipients []string    `json:"recipients"`
	Body       string      `json:"body"`
	Params     *sms.Params `json:"params"`
}
