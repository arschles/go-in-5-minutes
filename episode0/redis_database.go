package episode0

import (
	"fmt"
	"gopkg.in/redis.v3"
)

//RedisClientWrapper is a wrapper for the Client type the the gopkg.in/redis.v3 package
//This type allows to redefine the Get and Set method using the parameter types and return types that
//match the Hashtable interface.
type RedisClientWrapper struct {
	Client *redis.Client
}

//Get method of the RedisClientWrapper type
func (i RedisClientWrapper) Get(key string) ([]byte, error) {
	val, ok := i.Client.Get(key).Result()
	if ok != nil {
		return nil, fmt.Errorf("Error: %s", ok)
	}
	return []byte(val), nil
}

//Set method of the RedisClientWrapper type
func (i RedisClientWrapper) Set(key string, val []byte) error {
	err := i.Client.Set(key, val, 0).Err()
	return err
}
