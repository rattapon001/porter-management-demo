package job_kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConfing struct {
	ServerType string
	Host       string
	Port       string
}

func LoadKafkaConfig() *KafkaConfing {
	serverType := os.Getenv("KAFKA_SERVER_TYPE")
	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")

	return &KafkaConfing{
		ServerType: serverType,
		Host:       host,
		Port:       port,
	}

}

func CreateKafkaProducer() *kafka.Producer {
	config := LoadKafkaConfig()
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.Host + ":" + config.Port,
	}
	p, err := kafka.NewProducer(producerConfig)
	if err != nil {
		panic(err)
	}
	return p
}
