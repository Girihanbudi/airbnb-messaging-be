package usecase

import (
	"airbnb-messaging-be/internal/app/sms/preset/request"
	"airbnb-messaging-be/internal/app/sms/preset/response"
	"airbnb-messaging-be/internal/pkg/stderror"
	"context"
)

type ISms interface {
	SendSms(ctx context.Context, cmd request.SendSms) (res response.SendSms, err *stderror.StdError)
}
