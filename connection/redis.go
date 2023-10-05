package connection

import (
	"context"
	"log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {
	if config.CONFIG.REDIS_HOST != "" {
		redis := redis.NewClient(&redis.Options{
			Addr:     config.CONFIG.REDIS_HOST + ":" + config.CONFIG.REDIS_PORT,
			Password: config.CONFIG.REDIS_PASS,
			DB:       0,
		})

		_, err := redis.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		return redis, nil
	}

	log.Print("Redis is connected")

	return nil, nil
}

func Redis() *redis.Client {
	return connection.redis
}
