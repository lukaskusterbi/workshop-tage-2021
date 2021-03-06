package producer

import (
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaProducer struct {
	config   *Config
	name     string
	producer *kafka.Producer
	ticker   *time.Ticker
	stop     chan bool
	running  bool
	messages []string
}

type Config struct {
	kafkaBootstrapServers string
	kafkaPort             int
	kafkaTopic            string
}
