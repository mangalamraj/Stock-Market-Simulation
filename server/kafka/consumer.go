package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"stock-market-simulation/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redis/go-redis/v9"
)

// Stock represents a stock price update

func StartConsumer(brokers, groupID, topic string, rdb *redis.Client) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     brokers,
		"group.id":              groupID,
		"auto.offset.reset":     "earliest",
		"broker.address.family": "v4",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		var stock types.Stock
		err = json.Unmarshal(msg.Value, &stock)
		if err != nil {
			fmt.Printf("Error deserializing message: %v\n", err)
			continue
		}

		// Cache stock price in Redis
		// Cache stock price in Redis (for the latest price)
		err = rdb.Set(context.Background(), stock.Symbol, stock.Price, 0).Err()
		if err != nil {
			fmt.Printf("Error caching stock price: %v\n", err)
			continue
		}

		// Store historical stock price in Redis sorted set
		timestamp := float64(msg.Timestamp.Unix()) // Get timestamp from Kafka message
		err = rdb.ZAdd(context.Background(), stock.Symbol+":history", redis.Z{
			Score:  timestamp,
			Member: stock.Price,
		}).Err()
		if err != nil {
			fmt.Printf("Error storing historical price: %v\n", err)
			continue
		}

		fmt.Printf("Stored in Redis history: %s -> %.2f at %v\n", stock.Symbol, stock.Price, timestamp)

		fmt.Printf("Cached in Redis: %s -> %.2f\n", stock.Symbol, stock.Price)
	}
}
