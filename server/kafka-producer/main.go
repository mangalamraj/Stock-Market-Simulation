package main

import "stock-market-simulation/kafka"

func main() {
	brokers := "localhost:9092"
	topic := "stock-prices"
	kafka.StartProducer(brokers, topic)
}
