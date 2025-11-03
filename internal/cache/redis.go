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

func (r *Redis) SaveToken(ctx context.Context, token string, ttl time.Duration) error {
	return r.rClient.Set(ctx, token, "whitelist", ttl).Err()
}

func (r *Redis) GetToken(ctx context.Context, token string, ttl time.Duration) (string, error) {
	val, err := r.rClient.Get(ctx, token).Result()
	if err != nil {
		return "", err
	}
	return val, err

}

func (r *Redis) BanToken(ctx context.Context, token string, ttl time.Duration) error {
	return r.rClient.Set(ctx, token, "blacklist", ttl).Err()
}

func (r *Redis) IsTokenBanned(ctx context.Context, token string) bool {
	list, err := r.rClient.Get(ctx, token).Result()
	if err != nil {
		return false
	} else if list == "blacklist" {
		return true
	}
	return false
}
