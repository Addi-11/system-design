package main

import (
	"log"
	"github.com/Shopify/sarama"
)

func main() {
	// Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}
	defer producer.Close()

	// message to send to Kafka
	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello Kafka from Go!"),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message to Kafka: %s", err)
	} else {
		log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
	}
}
