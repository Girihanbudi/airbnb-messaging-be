package config

type Config struct {
	Host  string `mapstructure:"host"`
	Group string `mapstructure:"port"`
}
