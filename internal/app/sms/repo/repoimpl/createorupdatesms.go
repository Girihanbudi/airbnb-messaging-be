package repoimpl

import (
	module "airbnb-messaging-be/internal/app/sms"
	"context"
)

func (r Repo) CreateOrUpdateSms(ctx context.Context, sms *module.Sms) (err error) {
	err = r.Gorm.DB.Save(sms).Error
	return
}
