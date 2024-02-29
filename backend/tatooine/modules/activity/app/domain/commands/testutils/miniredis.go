package testutils

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

func SetupMiniredis(t *testing.T) (*miniredis.Miniredis, *cache.RedisClient) {
	t.Helper()

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' occurred when starting miniredis", err)
	}

	redisOpt := cache.RedisOption{
		Options: &redis.Options{
			Addr:     mr.Addr(),
			DB:       0,
			Password: "",
		},
		ExpirationTime: 1 * time.Hour,
	}
	client := cache.NewRedisClient(redisOpt)

	return mr, client
}
