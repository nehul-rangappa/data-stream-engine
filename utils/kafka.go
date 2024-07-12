package utils

import (
	"github.com/segmentio/kafka-go"
)

// CreateTopic takes a broker and topic and
// checks if the desired topic exists
// creates a new one if it doesn't exist
// and returns error if any
func CreateTopic(broker, topic string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return err
	}

	defer conn.Close()

	topics, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	isAvailable := false
	for _, p := range topics {
		if p.Topic == topic {
			isAvailable = true
			break
		}
	}

	if isAvailable {
		return nil
	}

	topicConfig := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfig...)
	if err != nil {
		return err
	}

	return nil
}
