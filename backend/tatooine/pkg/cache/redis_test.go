package cache_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestRedisClient_Set(t *testing.T) {
	mr, client := setupMiniredis(t)
	defer mr.Close()
	defer client.Close()

	ctx := context.Background()
	key := "testKey"
	value := "testValue"

	err := client.Set(ctx, key, value)

	assert.NoError(t, err)

}

func TestRedisClient_Get(t *testing.T) {
	mr, client := setupMiniredis(t)
	defer mr.Close()
	defer client.Close()

	ctx := context.Background()

	tests := []struct {
		name      string
		key       string
		value     string
		wantError bool
	}{
		{
			name:      "should get a key-value pair successfully",
			key:       "testKey",
			value:     "testValue",
			wantError: false,
		},

		{
			name:      "should return an error when key not exist",
			key:       "notExsistentKey",
			value:     "testValue",
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if !tc.wantError {
				client.Set(ctx, tc.key, tc.value)
			}

			result, err := client.Get(ctx, tc.key)

			if tc.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.value, result)
			}
		})
	}
}

func TestRedisClient_AddMember(t *testing.T) {
	mr, client := setupMiniredis(t)
	defer mr.Close()
	defer client.Close()

	ctx := context.Background()
	tests := []struct {
		name      string
		memberKey string
		members   []any
		wantError bool
	}{
		{
			name:      "should get a key-value pair successfully",
			memberKey: "memberKey",
			members:   []any{"member1", "member2"},
			wantError: false,
		},

		{
			name:      "should return an when members is nil",
			memberKey: "memberKey",
			members:   nil,
			wantError: true,
		},

		{
			name:      "should return an error when members is empty",
			memberKey: "memberKey",
			members:   []any{},
			wantError: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := client.AddMember(ctx, tc.memberKey, tc.members...)

			if tc.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRedisClient_GetMembers(t *testing.T) {
	mr, client := setupMiniredis(t)
	defer mr.Close()
	defer client.Close()

	ctx := context.Background()
	tests := []struct {
		name      string
		memberKey string
		members   []any
		wantError bool
	}{
		{
			name:      "should get members successfully",
			memberKey: "memberKey",
			members:   []any{"member1", "member2"},
			wantError: false,
		},

		{
			name:      "should return empty when memberKey is not exist",
			memberKey: "nonExistingMemberKey",
			members:   []any{"member1", "member2"},
			wantError: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if !tc.wantError {
				fmt.Println(tc.memberKey)
				client.AddMember(ctx, tc.memberKey, tc.members...)
			}

			result, err := client.GetMembers(ctx, tc.memberKey)

			if tc.wantError {
				assert.Empty(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result)
				assert.Equal(t, len(tc.members), len(result))
			}
		})
	}
}

func TestRedisClient_Delete(t *testing.T) {
	mr, client := setupMiniredis(t)
	defer mr.Close()
	defer client.Close()

	ctx := context.Background()

	key := "testKey"
	value := "testValue"

	client.Set(ctx, key, value)
	err := client.Delete(ctx, key)

	assert.NoError(t, err)

}

func setupMiniredis(t *testing.T) (*miniredis.Miniredis, *cache.RedisClient) {
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
