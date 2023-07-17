package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "books-kafka_kafka_1:9092",
		"client.id":         "books-consumer",
		"group.id":          "books-group",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("Error consumer", err.Error())
	}

	topics := []string{"bookMaster"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
