package main

import (
	"log"
	"github.com/Shopify/sarama"
)

func main() {
	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %s", err)
	}
	defer consumer.Close()
 
	// Subscribe to the topic 'test-topic'
	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume from Kafka: %s", err)
	}
	defer partitionConsumer.Close()

	// Read messages from Kafka
	for message := range partitionConsumer.Messages() {
		log.Printf("Message received: %s", string(message.Value))
	}
}
