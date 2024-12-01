package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func PublishMessage(topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Fatalf("Failed to write messages to Kafka: %v", err)
	}
	log.Println("Message published to Kafka:", message)
}
