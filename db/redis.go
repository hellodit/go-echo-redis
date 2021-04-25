package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Client struct {
	*redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

var once sync.Once
var client *Client

func RedisClient() {
	HOST := viper.GetString("REDIS_HOST")
	PORT := viper.GetString("REDIS_PORT")
	PASS := viper.GetString("REDIS_PASSWORD")

	once.Do(func() {
		conn := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", HOST, PORT),
			Password: PASS,
			DB:       0,
		})

		_, err := conn.Ping().Result()
		if err != nil {
			logrus.Fatalf("Could not connect to redis %v", err)
		}

		client = &Client{conn}
	})
	logrus.Info("Success connect to Redis")
}

func GetRedis() *Client {
	if client == nil {
		RedisClient()
	}

	return client
}
