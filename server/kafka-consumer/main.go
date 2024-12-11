package main

import (
	"stock-market-simulation/kafka"

	"github.com/redis/go-redis/v9"
)

func main() {
	brokers := "localhost:9092"
	groupID := "stock-consumer-group"
	topic := "stock-prices"

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	kafka.StartConsumer(brokers, groupID, topic, rdb)
}
