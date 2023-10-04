package connection

import (
	"context"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedisGatewatch() (*redis.Client, error) {
	if config.CONFIG.REDIS_HOST != "" {
		redis := redis.NewClient(&redis.Options{
			Addr:     config.CONFIG.REDIS_HOST + ":" + config.CONFIG.REDIS_PORT,
			Password: config.CONFIG.REDIS_PASS,
			DB:       0,
		})

		_, err := redis.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}

		return redis, nil
	}

	return nil, nil
}

func RedisGatewatch() *redis.Client {
	return database.redis
}
