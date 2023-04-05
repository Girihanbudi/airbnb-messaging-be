package repo

import (
	module "airbnb-messaging-be/internal/app/sms"
	"context"
)

type ISms interface {
	CreateOrUpdateSms(ctx context.Context, sms *module.Sms) (err error)
}
