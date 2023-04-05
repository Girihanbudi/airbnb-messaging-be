package usecaseimpl

import (
	module "airbnb-messaging-be/internal/app/sms"
	"airbnb-messaging-be/internal/app/sms/preset/request"
	"airbnb-messaging-be/internal/app/sms/preset/response"
	"airbnb-messaging-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) SendSms(ctx context.Context, req request.SendSms) (res response.SendSms, err *stderror.StdError) {
	sms := module.Sms{
		Type:    req.Type,
		Context: req.Context,
		Payload: req.Payload,
	}

	if valid, validateErr := req.Validate(); !valid {
		sms.Status = module.SmsStatusFailed.String()
		if err = u.CreateSmsLog(ctx, &sms); err != nil {
			return
		}
		err = stderror.DEF_DATA_400.ErrorMsg(validateErr)
		return
	}

	var payload module.Payload
	if extractPayloadErr := req.Payload.Scan(&payload); err != nil {
		sms.Status = module.SmsStatusFailed.String()
		if err = u.CreateSmsLog(ctx, &sms); err != nil {
			return
		}
		err = stderror.DEF_DATA_400.ErrorMsg(extractPayloadErr)
		return
	}

	_, sendErr := u.Messenger.SendSms(payload.Recipients, payload.Body, nil, payload.Params)
	if sendErr != nil {
		sms.Status = module.SmsStatusFailed.String()
		if err = u.CreateSmsLog(ctx, &sms); err != nil {
			return
		}
		err = stderror.DEF_SERVER_503.ErrorMsg(sendErr)
		return
	}

	sms.Status = module.SmsStatusSent.String()
	err = u.CreateSmsLog(ctx, &sms)
	res.SmsId = &sms.Id

	return
}

func (u Usecase) CreateSmsLog(ctx context.Context, sms *module.Sms) (err *stderror.StdError) {
	if createErr := u.SmsRepo.CreateOrUpdateSms(ctx, sms); createErr != nil {
		err = stderror.DEF_SERVER_503.ErrorMsg(createErr)
		return
	}

	return
}
