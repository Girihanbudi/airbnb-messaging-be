package event

func (h Handler) RegisterApi() {
	sms := h.Router.Group("sms")
	{
		send := sms.Group("send")
		{
			send.Listen("init", h.SendSms)
			// send.Listen("success", h.SendSms)
			// send.Listen("failed", h.SendSms)
		}
	}
}
