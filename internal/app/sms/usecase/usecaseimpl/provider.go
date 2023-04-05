package usecaseimpl

import (
	"airbnb-messaging-be/internal/app/sms/repo"

	"airbnb-messaging-be/internal/pkg/messagebird"
)

type Options struct {
	SmsRepo   repo.ISms
	Messenger *messagebird.Messenger
}

type Usecase struct {
	Options
}

func NewSmsUsecase(options Options) *Usecase {
	return &Usecase{options}
}
