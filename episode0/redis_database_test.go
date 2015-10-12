package episode0

import (
	"gopkg.in/redis.v3"
	"testing"
)

//Test our BusinessLogic using our RedisClientWrapper
//You need to have a redis-server running to pass this test.
//(Note that it is way better to use the in_memory_hashtable to test our BusinessLogic.)
//command: cd episode0 && go test *.go
func TestBusinessLogicWithRedis(t *testing.T) {
	//we create our RedisClientWrapper that match the hashtable interface

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
