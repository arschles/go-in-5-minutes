package episode0

import (
	"fmt"
	"gopkg.in/redis.v3"
)

type RedisClientWrapper struct {
	Client *redis.Client
}

func (i RedisClientWrapper) Get(key string) ([]byte, error) {
	val, ok := i.Client.Get(key).Result()
	if ok != nil {
		return nil, fmt.Errorf("Error: %s", ok)
	}
	return []byte(val), nil
}

func (i RedisClientWrapper) Set(key string, val []byte) error {
	err := i.Client.Set(key, val, 0).Err()
	return err
}
