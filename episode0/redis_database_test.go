package episode0

import (
	"gopkg.in/redis.v3"
	"testing"
)

func TestBusinessLogicWithRedis(t *testing.T) {
	//in memory hash table is a mock
	ht := RedisClientWrapper{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
	BusinessLogic(ht)
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got '%s'", string(val))
	}
}
