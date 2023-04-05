package config

import (
	cache "airbnb-messaging-be/internal/pkg/cache/config"
	gorm "airbnb-messaging-be/internal/pkg/gorm/config"
	httpserver "airbnb-messaging-be/internal/pkg/http/server/config"
	jwt "airbnb-messaging-be/internal/pkg/jwt/config"
	kafka "airbnb-messaging-be/internal/pkg/kafka/config"
	messagebird "airbnb-messaging-be/internal/pkg/messagebird/config"
)

type Config struct {
	Stage      string             `mapstructure:"stage"`
	Origins    []string           `mapstructure:"origins"`
	Domain     string             `mapstructure:"domain"`
	HttpServer httpserver.Config  `mapstructure:"httpserver"`
	DB         gorm.Config        `mapstructure:"db"`
	Jwt        jwt.Config         `mapstructure:"jwt"`
	Cache      cache.Config       `mapstructure:"cache"`
	Messenger  messagebird.Config `mapstructure:"messenger"`
	Kafka      kafka.Config       `mapstructure:"kafka"`
}
