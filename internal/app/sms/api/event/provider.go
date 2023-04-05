package event

import (
	smsusecase "airbnb-messaging-be/internal/app/sms/usecase"
	kafka "airbnb-messaging-be/internal/pkg/kafka/router"
)

type Options struct {
	Sms    smsusecase.ISms
	Router *kafka.Router
}

type Handler struct {
	Options
}

func NewSmsHandler(options Options) *Handler {
	return &Handler{options}
}
