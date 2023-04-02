package config

type Config struct {
	AccessKey  string `mapstructure:"accesskey"`
	Originator string `mapstructure:"originator"`
}
