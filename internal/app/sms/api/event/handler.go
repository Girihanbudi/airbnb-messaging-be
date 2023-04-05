package event

import (
	"airbnb-messaging-be/internal/app/sms/preset/request"
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (h Handler) SendSms(ctx context.Context, msg *kafka.Message) {
	var req request.SendSms
	json.Unmarshal(msg.Value, &req)
	h.Sms.SendSms(ctx, req)
}
