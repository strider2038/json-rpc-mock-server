package app

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Protocol    string `default:"tcp"`
	BearerToken string `split_words:"true"`
	Port        uint16 `default:"4000"`
}

func LoadConfigFromEnvironment() Config {
	config := Config{}
	envconfig.MustProcess("JSON_RPC", &config)

	return config
}
