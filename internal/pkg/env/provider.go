package env

import (
	"airbnb-messaging-be/internal/pkg/env/config"
	"airbnb-messaging-be/internal/pkg/log"
	"fmt"

	"github.com/spf13/viper"
)

const Instance string = "Env"

// global env declaration
var CONFIG config.Config

type EnvConfig struct {
	Path     string
	FileName string
	Ext      string
}

func NewDefaultEnvConf() EnvConfig {
	return EnvConfig{
		Path:     "./env",
		FileName: "config",
		Ext:      "yaml",
	}
}

func InitEnv(conf EnvConfig) {
	log.Event(Instance, "reading config...")

	viper.AddConfigPath(conf.Path)
	viper.SetConfigName(conf.FileName)
	viper.SetConfigType(conf.Ext)

	env := config.Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(Instance, "failed to read config", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(Instance, "failed to unmarshal config", err)
	}

	log.Event(Instance, fmt.Sprintf("using %s stage mode", env.Stage))
	CONFIG = env
}

func ProvideEnv() config.Config {
	return CONFIG
}