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

func NewRedis() *Redis {
	return &Redis{}
}

//Get method of the RedisClientWrapper type
func (r *Redis) Save(m models.Model) error {
	b, err := m.MarshalBinary()
	if err != nil {
		return err
	}
	return r.client.Set(m.Key().String(), b, time.Duration(0)).Err()
}

func (r *Redis) Delete(m models.Model) error {
	return r.client.Del(m.Key().String()).Err()
}

func (r *Redis) Get(pk models.PrimaryKey, m models.Model) error {
	str, err := r.client.Get(pk.String()).Result()
	if err != nil {
		return err
	}
	return m.UnmarshalBinary([]byte(str))
}
