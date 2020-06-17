package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Configuration struct {
	Port int    `env:"Port" envDefault:"3000"`
	Home string `env:"Home"`
}

func NewConfiguration() Configuration {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Errorf("%+v\n", err)
	}
	return cfg
}
