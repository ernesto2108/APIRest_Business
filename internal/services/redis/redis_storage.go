package internal

import (
	"context"
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RdsClient struct {
	client *redis.Client
}

func getCon() *RdsClient {
	logs.Log().Info("connecting redis")
	rdb, err := redis.ParseURL("redis://:qwerty@localhost:6379")
	if err != nil {
		logs.Log().Fatal("Cannot connect with Redis")
		panic(err)
	}

	client := redis.NewClient(rdb)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		logs.Log().Fatal("Failed to ping Redis")
		panic(err)
	}

	return &RdsClient{client}
}
