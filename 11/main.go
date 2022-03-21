package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "xxx",
		Password: "xxx",
		DB:       0,
		PoolSize: 10000,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("initialize redis failed, err: %v\n", err)
	} else {
		fmt.Printf("initialize redis successfully\n")
	}
	pubsub := rdb.Subscribe(ctx, "ch1")
	defer pubsub.Close()
	ch := pubsub.Channel()

	err = rdb.Publish(ctx, "ch1", "发布消息").Err()
	if err != nil {
		panic(err)
	}
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
