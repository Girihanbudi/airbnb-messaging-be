package tool

import (
	"airbnb-messaging-be/internal/pkg/env/config"
	gorm "airbnb-messaging-be/internal/pkg/gorm/config"
	httpServer "airbnb-messaging-be/internal/pkg/http/server/config"
)

func ExtractServerConfig(config config.Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config config.Config) gorm.Config {
	return config.DB
}
