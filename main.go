package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nehul-rangappa/data-stream-engine/handler"
	"github.com/nehul-rangappa/data-stream-engine/service"
	"github.com/nehul-rangappa/data-stream-engine/utils"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Load environment variables from config file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
	}

	// Read environment variables
	broker := os.Getenv("BOOTSTRAP_SERVER")
	consumerTopic := os.Getenv("KAFKA_CONSUMER_TOPIC")
	producerTopic := os.Getenv("KAFKA_PRODUCER_TOPIC")

	// Create a new topic for storing processed data if it doesn't exists
	err = utils.CreateTopic(broker, producerTopic)
	if err != nil {
		log.Fatal("failed to create topic:", err)
	}

	// Initialize Kafka consumer and producer
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   consumerTopic,
	})

	defer consumer.Close()

	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   producerTopic,
	})

	defer producer.Close()

	kafkaService := service.New()
	kafkaHandler := handler.New(kafkaService, consumer, producer)

	kafkaHandler.GetUserData(context.Background())
}
