package redisdb

import (
	"crypto/tls"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/L1LSunflower/auction/config"
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
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Address, cfg.Redis.Port),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: cfg.Redis.Address,
		},
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
