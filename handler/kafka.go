package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nehul-rangappa/data-stream-engine/model"
	"github.com/nehul-rangappa/data-stream-engine/service"
	"github.com/segmentio/kafka-go"
)

// KafkaHandler defines the attributes needed for handler
type KafkaHandler struct {
	service  *service.DataInsights
	consumer *kafka.Reader
	producer *kafka.Writer
}

// New is the factory function for handler layer
func New(service *service.DataInsights, consumer *kafka.Reader, producer *kafka.Writer) *KafkaHandler {
	return &KafkaHandler{
		service:  service,
		consumer: consumer,
		producer: producer,
	}
}

// GetUserData method takes a context
// continuously listens to the data stream for events
// processes every message and produces it to another topic
func (h *KafkaHandler) GetUserData(ctx context.Context) {
	for {
		var msg model.Message

		m, err := h.consumer.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		// log.Printf("Received message: %s\n", string(m.Value))

		if err := json.Unmarshal(m.Value, &msg); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		processedMessage := h.service.ProcessData(msg)
		if processedMessage == nil {
			continue
		}

		log.Printf("Processed and filtered message: %v", string(processedMessage))

		// Forward the processed message to the next topic
		err = h.producer.WriteMessages(ctx, kafka.Message{
			Value: processedMessage,
		})
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		}
	}
}
