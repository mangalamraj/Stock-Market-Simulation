package kafka

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"stock-market-simulation/types"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartProducer(brokers, topic string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":     brokers,
		"broker.address.family": "v4",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	stocks := []string{"AAPL", "GOOG", "TSLA", "AMZN"}

	for {
		for _, stock := range stocks {
			price := rand.Float64()*1000 + 100 // Random price between $100-$1100
			stockData := types.Stock{Symbol: stock, Price: price}

			message, err := json.Marshal(stockData)
			if err != nil {
				fmt.Printf("Failed to serialize message: %v\n", err)
				continue
			}

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

		time.Sleep(1 * time.Second)
	}
}

func BuyProducer(brokers, topic string, stock string, price float64) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":     brokers,
		"broker.address.family": "v4",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	stockData := types.Stock{Symbol: stock, Price: price}

	message, err := json.Marshal(stockData)
	if err != nil {
		fmt.Printf("Failed to serialize message: %v\n", err)
		return
	}
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
