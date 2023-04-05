package request

import (
	"airbnb-messaging-be/internal/pkg/json"
	"airbnb-messaging-be/internal/pkg/validator"
)

type SendSms struct {
	Type    string   `json:"type" validator:"required"`
	Context string   `json:"context" validator:"required"`
	Payload json.Raw `json:"payload" validator:"required"`
}

func (req *SendSms) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
