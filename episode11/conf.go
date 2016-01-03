package main

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "episode11"
)

type Config struct {
	Port int `envconfig:"port" default:"8080"`
}

func GetConfig() (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(AppName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
