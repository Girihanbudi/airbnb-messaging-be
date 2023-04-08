package event

import (
	"airbnb-messaging-be/internal/app/sms/preset/request"
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
)

func (h Handler) SendSms(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req request.SendSms
	json.Unmarshal(msg.Value, &req)
	h.Sms.SendSms(ctx, req)
}
