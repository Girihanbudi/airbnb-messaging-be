package config

import (
	router "airbnb-messaging-be/internal/pkg/kafka/router/config"
)

type Config struct {
	Host   string        `mapstructure:"host"`
	Group  string        `mapstructure:"port"`
	Router router.Config `mapstructure:"router"`
}
