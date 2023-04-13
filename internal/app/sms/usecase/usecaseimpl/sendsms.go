package usecaseimpl

import (
	module "airbnb-messaging-be/internal/app/sms"
	"airbnb-messaging-be/internal/app/sms/preset/request"
	"airbnb-messaging-be/internal/app/sms/preset/response"
	"airbnb-messaging-be/internal/pkg/stderror"
	"context"
	"fmt"
)

func (u Usecase) SendSms(ctx context.Context, cmd request.SendSms) (res response.SendSms, err *stderror.StdError) {
	sms := module.Sms{
		Type:    cmd.Type,
		Context: cmd.Context,
		Payload: cmd.Payload,
	}

	if valid, validateErr := cmd.Validate(); !valid {
		sms.Status = module.SmsStatusFailed.String()
		sms.Logs = fmt.Sprintf("failed to validate request, %s", validateErr)
		if err = u.CreateSmsLog(ctx, &sms); err != nil {
			return
		}
		err = stderror.DEF_DATA_400.ErrorMsg(validateErr)
		return
	}

	var payload module.Payload
	if extractPayloadErr := cmd.Payload.Get(&payload); extractPayloadErr != nil {
		sms.Status = module.SmsStatusFailed.String()
		sms.Logs = fmt.Sprintf("failed to extract request payload, %s", extractPayloadErr)
		if err = u.CreateSmsLog(ctx, &sms); err != nil {
			return
		}
		err = stderror.DEF_DATA_400.ErrorMsg(extractPayloadErr)
		return
	}

	if _, sendErr := u.Messenger.SendSms(payload.Recipients, payload.Body, nil, payload.Params); sendErr != nil {
		sms.Status = module.SmsStatusFailed.String()
		sms.Logs = fmt.Sprintf("failed to send sms to provider, %s", sendErr)
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
