package kafka

import (
	"airbnb-messaging-be/internal/pkg/kafka/config"
	"airbnb-messaging-be/internal/pkg/kafka/router"
	"airbnb-messaging-be/internal/pkg/log"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const Instance string = "Kafka Consumer"

type TopicHandler struct {
	Topic   string
	Handler func()
}

type Options struct {
	config.Config
	Router *router.Router
}

type Listener struct {
	isListening bool

	Consumer *kafka.Consumer
	Options
}

func NewEventListener(options Options) *Listener {

	newConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": options.Host,
		"group.id":          options.Group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(Instance, "failed to init event listener", err)
	}

	log.Event(Instance, fmt.Sprintf("ready to consume from %s", options.Host))

	return &Listener{
		Consumer: newConsumer,
		Options:  options,
	}
}
