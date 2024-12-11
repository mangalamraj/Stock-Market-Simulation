package kafka

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"stock-market-simulation/types"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Stock represents a stock price update


func StartProducer(brokers, topic string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"broker.address.family": "v4",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	stocks := []string{"AAPL", "GOOG", "TSLA", "AMZN"}

	for {
		for _, stock := range stocks {
			// Simulate stock price
			price := rand.Float64()*1000 + 100 // Random price between $100-$1100
			stockData := types.Stock{Symbol: stock, Price: price}

			// Serialize to JSON
			message, err := json.Marshal(stockData)
			if err != nil {
				fmt.Printf("Failed to serialize message: %v\n", err)
				continue
			}

			// Produce message
			err = producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          message,
			}, nil)

			if err != nil {
				fmt.Printf("Failed to produce message: %v\n", err)
			} else {
				fmt.Printf("Produced: %s -> %.2f\n", stock, price)
			}
		}

		// Wait 1 second between batches
		time.Sleep(1 * time.Second)
	}
}
