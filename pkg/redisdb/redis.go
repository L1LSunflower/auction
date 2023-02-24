package redisdb

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

type RedisConnection struct {
	RedisClient *redis.Client
}

var onceRedisConnectionInstance sync.Once
var RedisConnectionInstance *RedisConnection

func RedisInstance() *RedisConnection {

	if RedisConnectionInstance == nil {
		onceRedisConnectionInstance.Do(func() {
			conn, err := redisConnect()

			if err != nil {
				log.Println("redisdb connect error:", err)
			}

			RedisConnectionInstance = conn
		})
	} else {
		err := RedisConnectionInstance.Ping()

		if err != nil {
			log.Println("redisdb ping error:", err)
		}
	}

	return RedisConnectionInstance
}

func redisConnect() (*RedisConnection, error) {
	cfg := config.GetConfig()
	logger.Log.Info(message.NewMessage(fmt.Sprintf("Redis address: %s, Redis Port: %s, Redis Password: %s\n", cfg.Redis.Address, cfg.Redis.Port, cfg.Redis.Password)))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Address, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping(client.Context()).Result()

	if err != nil {
		return nil, err
	}

	return &RedisConnection{
		RedisClient: client,
	}, nil
}

func (r *RedisConnection) Ping() error {
	_, err := r.RedisClient.Ping(r.RedisClient.Context()).Result()

	if err == nil {
		return nil
	}

	RedisConnection, err := redisConnect()

	if err != nil {
		return err
	}

	_, err = RedisConnection.RedisClient.Ping(RedisConnection.RedisClient.Context()).Result()

	if err != nil {
		return err
	}

	return nil
}
