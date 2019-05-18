package redis_mock

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

func NewMockRedis(t *testing.T) *redis.Client {
	t.Helper()

	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("unexpected error while creagtign test redis server '%#v'", err)
	}

	// create *redis.Client
	client := redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "",
		DB:       0,
	})
	return client
}
