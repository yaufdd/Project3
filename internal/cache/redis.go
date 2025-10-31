package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Redis {
	return &Redis{rdb: rdb}
}

func (r *Redis) IsTokenBanned(ctx context.Context, token string) bool {
	key := "blacklist:" + token
	exist, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return exist == 1
}

func (r *Redis) BanToken(ctx context.Context, token string, ttl time.Duration) error {
	key := "blacklist:" + token
	return r.rdb.Set(ctx, key, 1, ttl).Err()
}
