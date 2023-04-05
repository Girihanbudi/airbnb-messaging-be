package kafka

import (
	"airbnb-messaging-be/internal/pkg/kafka/router"
	"airbnb-messaging-be/internal/pkg/log"
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/thoas/go-funk"
)

func (l *Listener) Start() error {
	topics := funk.Map(l.Router.Handlers, func(handler router.Handler) string {
		return handler.Topic
	}).([]string)

	l.Consumer.SubscribeTopics(topics, nil)

	log.Event(Instance, "starting listener...")
	l.isListening = true

	for l.isListening {
		msg, err := l.Consumer.ReadMessage(time.Second)
		for _, handler := range l.Router.Handlers {
			if handler.Topic == string(*msg.TopicPartition.Topic) {
				handler.Handler(context.Background(), msg)
			}
		}
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			log.Error(Instance, fmt.Sprintf("consumer error: %v", msg), err)
		} else {
			return err
		}
	}

	return nil
}

func (l *Listener) Stop() error {
	log.Event(Instance, "shutting down listener...")
	l.isListening = false
	l.Consumer.Close()

	return nil
}
