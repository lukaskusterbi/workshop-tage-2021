package consumer

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaConsumer struct {
	config   *Config
	name     string
	consumer *kafka.Consumer
	running  bool
}

type Config struct {
	kafkaBootstrapServers string
	kafkaPort             int
	kafkaTopics           []string
	kafkaConsumerGroup    string
	kafkaAutoOffsetReset  string
}
