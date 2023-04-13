package sms

import (
	"airbnb-messaging-be/internal/pkg/json"
	"time"
)

type Sms struct {
	Id      uint     `json:"id" gorm:"primaryKey"`
	Type    string   `json:"type"`    //message type (info, secrect, warning, etc.)
	Context string   `json:"context"` //what is the purpose of this message (login authorization, info, etc.)
	Payload json.Raw `json:"payload"` //message payload sent
	Status  string   `json:"status"`  //message status after send
	Logs    string   `json:"logs"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
