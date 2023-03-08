package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, servers string, msgChan chan *ckafka.Message) {
	kafkaConsumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "goapp",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	err = kafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Kafka consumer has started")

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
