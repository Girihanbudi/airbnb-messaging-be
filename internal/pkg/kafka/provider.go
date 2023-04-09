package kafka

import (
	"airbnb-messaging-be/internal/pkg/kafka/config"
	"airbnb-messaging-be/internal/pkg/kafka/router"
	"airbnb-messaging-be/internal/pkg/log"
	"fmt"

	"github.com/Shopify/sarama"
)

const Instance string = "Kafka Client"

type TopicHandler struct {
	Topic   string
	Handler func()
}

type Options struct {
	config.Config
	Router *router.Router
}

type Listener struct {
	isReady chan bool

	Consumer sarama.ConsumerGroup
	Options
}

func NewEventListener(options Options) *Listener {

	sarama.Logger = log.NewLogger(Instance, false)
	version, err := sarama.ParseKafkaVersion(options.Version)
	if err != nil {
		log.Fatal(Instance, "error parsing Kafka version", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version
	config.ClientID = options.ClientId

	switch options.Assigner {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Fatal(Instance, fmt.Sprintf("unrecognized consumer group partition assignor: %s", options.Assigner), nil)
	}

	if options.IsUseOldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	client, err := sarama.NewConsumerGroup(options.Brokers, options.Group, config)
	if err != nil {
		log.Fatal(Instance, "error creating consumer group client: %v", err)
	}

	log.Event(Instance, fmt.Sprintf("ready to consume from %v", options.Brokers))

	return &Listener{
		isReady:  make(chan bool),
		Consumer: client,
		Options:  options,
	}
}
