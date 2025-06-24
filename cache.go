package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})
}

func GetFromCache(city string) (string, bool) {
	val, err := rdb.Get(ctx, city).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func SetCache(city, data string) {
	rdb.Set(ctx, city, data, 10*time.Minute)
}
