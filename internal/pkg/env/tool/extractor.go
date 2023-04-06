package tool

import (
	"airbnb-messaging-be/internal/pkg/env/config"
	gorm "airbnb-messaging-be/internal/pkg/gorm/config"
	httpserver "airbnb-messaging-be/internal/pkg/http/server/config"
	kafka "airbnb-messaging-be/internal/pkg/kafka/config"
	kafkarouter "airbnb-messaging-be/internal/pkg/kafka/router/config"
	messagebird "airbnb-messaging-be/internal/pkg/messagebird/config"
)

func ExtractServerConfig(config config.Config) httpserver.Config {
	return config.HttpServer
}

func ExtractDBConfig(config config.Config) gorm.Config {
	return config.DB
}

func ExtractMessengerConfig(config config.Config) messagebird.Config {
	return config.Messenger
}

func ExtractKafkaConfig(config config.Config) kafka.Config {
	return config.Kafka
}

func ExtractKafkaRouterConfig(config config.Config) kafkarouter.Config {
	return config.Kafka.Router
}
