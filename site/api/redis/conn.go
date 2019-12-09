package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arschles/go-in-5-minutes/site/api/models"
	"github.com/go-redis/redis"
	"github.com/gobuffalo/envy"
)

const castsHashKey = "screencasts"

func New(ctx context.Context) *redis.Client {
	redisHost := envy.Get("REDIS_HOST", "localhost")
	redisPort := envy.Get("REDIS_PORT", "6379")
	redisPass := envy.Get("REDIS_PASSWORD", "")
	redisDB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPass,
		DB:       0,
	})
	return redisDB
}

func AllEpisodes(ctx context.Context, cl *redis.Client) ([]models.Screencast, error) {
	scanCmd := cl.ZScan(castsHashKey, 0, "", 0)
	cl.Process(scanCmd)
	iter := scanCmd.Iterator()
	ret := []models.Screencast{}
	for iter.Next() {
		val := iter.Val()
		cast := models.Screencast{}
		if err := json.Unmarshal([]byte(val), &cast); err != nil {
			log.Printf("Error unmarshaling screencast (see JSON below)")
			log.Printf(val)
		}
		ret = append(ret, cast)
	}
	if err := iter.Err(); err != nil {
		log.Printf("Error iterating! %s", err)
		return nil, err
	}
	return ret, nil
}

func AddEpisode(ctx context.Context, cl *redis.Client, cast models.Screencast) error {
	// the score of the screencast in the sorted set is the unix timestamp
	score := cast.Date.Unix()
	elt := redis.Z{
		Score:  float64(score),
		Member: cast,
	}
	addCmd := cl.ZAdd(castsHashKey, elt)
	cl.Process(addCmd)
	if _, err := addCmd.Result(); err != nil {
		log.Printf("Error adding screencast! %s", err)
		return err
	}
	return nil
}
