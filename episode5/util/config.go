package util

import "fmt"

type Env string

func (e Env) String() string {
	return string(e)
}

const (
	EnvDev   Env = "dev"
	EnvStage Env = "stage"
	EnvProd  Env = "prod"
)

type Config struct {
	Port        int    `envconfig:"port"`
	Environment string `envconfig:"environment"`
	MongoURL    string `envconfig:"mongo_url"`
}

func (c Config) Env() (Env, error) {
	switch c.Environment {
	case EnvDev.String():
		return EnvDev, nil
	case EnvStage.String():
		return EnvStage, nil
	case EnvProd.String():
		return EnvProd, nil
	default:
		return Env(""), fmt.Errorf("invalid environment %s", c.Environment)
	}
}
