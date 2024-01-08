package msg_relay

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rattapon001/porter-management-demo/internal/job/domain"
	job_kafka "github.com/rattapon001/porter-management-demo/internal/job/infra/kafka"
	job_mongo "github.com/rattapon001/porter-management-demo/internal/job/infra/mongo"
)

type MyObject struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func MsgRelay(ctx context.Context) {
	producer := job_kafka.CreateKafkaProducer()
	db := job_mongo.MongoDbInit()
	changeStream := ChangeCapture(db, "porter_management", "jobs", ctx)
	for changeStream.Next(ctx) {
		var changeEvent struct {
			FullDocument *domain.Job `bson:"fullDocument"`
		}
		if err := changeStream.Decode(&changeEvent); err != nil {
			log.Println("Error decoding change event:", err)
			continue
		}
		// Handle the change event as needed
		if changeEvent.FullDocument != nil {
			// Handle the change event as needed
			fmt.Printf("Job: %+v\n", *changeEvent.FullDocument)
			data := changeEvent.FullDocument.Event[len(changeEvent.FullDocument.Event)-1]
			serializedObj, err := json.Marshal(data)

			topic := "your_topic"
			message := &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          serializedObj,
			}
			deliveryChan := make(chan kafka.Event)
			// Produce the message
			err = producer.Produce(message, deliveryChan)
			if err != nil {
				fmt.Printf("Error producing message: %v\n", err)
				return
			}

			e := <-deliveryChan
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		}
	}
	defer producer.Close()
}
