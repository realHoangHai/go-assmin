package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/realHoangHai/go-assmin/config"
	"strings"
	"time"
)

// ProviderRedisSet is redis providers.
var ProviderRedisSet = wire.NewSet(NewRedis)

// NewRedis is new redis.
func NewRedis(ctx context.Context) (redis.UniversalClient, error) {
	addrs := strings.Split(strings.Trim(config.C.Redis.Address, " "), ",")

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    addrs,
		Password: config.C.Redis.Password, // no password set
		DB:       config.C.Redis.DB,       // use default DB

		PoolSize: 1000,
	})

	cwt, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := rdb.Ping(cwt).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
