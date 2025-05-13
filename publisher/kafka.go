package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"iotApp/sensors" // Replace with your actual module name

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer *kafka.Writer
	topic  string
}

func ensureTopic(broker, topic string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	// Check if topic exists
	partitions, err := conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("failed to read partitions: %w", err)
	}
	for _, p := range partitions {
		if p.Topic == topic {
			return nil // topic exists
		}
	}

	// Create topic if it doesn't exist
	topicConfigs := []kafka.TopicConfig{{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}}
	return conn.CreateTopics(topicConfigs...)
}

func NewKafkaPublisher(brokerAddress, topic string) *KafkaPublisher {
	err := ensureTopic(brokerAddress, topic)
	if err != nil {
		panic(fmt.Sprintf("could not ensure topic exists: %v", err))
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &KafkaPublisher{
		writer: writer,
		topic:  topic,
	}
}

func (k *KafkaPublisher) Publish(data sensors.SensorStruct) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(data.SensorID),
		Value: payload,
		Time:  time.Now(),
	}

	return k.writer.WriteMessages(context.Background(), msg)
}
