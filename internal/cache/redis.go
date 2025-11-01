package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rClient *redis.Client
}

func NewCache(rdb *redis.Client) *Redis {
	return &Redis{rClient: rdb}
}

// "+" - whitelist
// "-" - blacklist
func (r *Redis) SaveToken(ctx context.Context, tokenHash string, ttl time.Duration) error {
	return r.rClient.Set(ctx, tokenHash, "+", ttl).Err()
}

func (r *Redis) BanToken(ctx context.Context, tokenHash string, ttl time.Duration) error {
	return r.rClient.Set(ctx, tokenHash, "-", ttl).Err()
}

func (r *Redis) IsTokenBanned(ctx context.Context, tokenHash string) bool {
	list, err := r.rClient.Get(ctx, tokenHash).Result()
	if err != nil {
		return false
	} else if list == "+" {
		return true
	}
	return false
}
