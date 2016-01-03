package main

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "episode11"
)

// Config is the envconfig-compatible configuration struct for this server. See https://github.com/kelseyhightower/envconfig for more detail
type Config struct {
	Port      int    `envconfig:"port" default:"8080"`
	DBType    string `envconfig:"db_type" default:"mem"`
	RedisHost string `envconfig:"redis_host" default:"localhost:6379"`
	RedisPass string `envconfig:"redis_pass" default:""` // default to no password
	RedisDB   int64  `envconfig:"redis_db" default:"0"`  // default to the redis default DB
}

// GetConfig uses envconfig to populate and return a Config struct. Returns all envconfig errors if they occurred
func GetConfig() (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(AppName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
