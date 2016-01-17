package db

import (
	"time"

	"github.com/arschles/go-in-5-minutes/episode11/models"
	"gopkg.in/redis.v3"
)

//RedisRedis is a db.DB implementation that talks to Redis
type Redis struct {
	client *redis.Client
}

// NewRedis initializes and returns a redis DB using the given raw redis.v3 client
func NewRedis(cl *redis.Client) *Redis {
	return &Redis{client: cl}
}

// Save is the interface implementation
func (r *Redis) Save(k models.Key, m models.Model) error {
	b, err := m.MarshalBinary()
	if err != nil {
		return err
	}
	return r.client.Set(k.String(), b, time.Duration(0)).Err()
}

// Delete is the interface implementation
func (r *Redis) Delete(k models.Key) error {
	return r.client.Del(k.String()).Err()
}

// Get is the interface implementation
func (r *Redis) Get(k models.Key, m models.Model) error {
	str, err := r.client.Get(k.String()).Result()
	if err != nil {
		return err
	}
	return m.UnmarshalBinary([]byte(str))
}
