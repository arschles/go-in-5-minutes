package main

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "episode11"
)

type Config struct {
	Port      int    `envconfig:"port" default:"8080"`
	RedisHost string `envconfig:"redis_host" default:"localhost:6379"`
	RedisPass string `envconfig:"redis_pass" default:""` // default to no password
	RedisDB   int64  `envconfig:"redis_db" default:"0"`  // default to the redis default DB
}

func GetConfig() (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(AppName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
