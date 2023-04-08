package config

import (
	router "airbnb-messaging-be/internal/pkg/kafka/router/config"
)

type Config struct {
	ClientId    string        `mapstructure:"clientid"`
	Brokers     []string      `mapstructure:"brokers"`
	Group       string        `mapstructure:"group"`
	Assigner    string        `mapstructure:"assigner"`
	Version     string        `mapstructure:"version"`
	IsUseOldest bool          `mapstructure:"isuseoldest"`
	Router      router.Config `mapstructure:"router"`
}
