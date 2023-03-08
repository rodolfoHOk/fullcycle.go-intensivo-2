package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rodolfoHOk/fullcycle.go-intensivo-2/project-1/internal/infra/database"
	"github.com/rodolfoHOk/fullcycle.go-intensivo-2/project-1/internal/usecase"
	"github.com/rodolfoHOk/fullcycle.go-intensivo-2/project-1/pkg/kafka"
	"github.com/rodolfoHOk/fullcycle.go-intensivo-2/project-1/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	// Kafka Consumer
	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	go kafka.Consume(topics, servers, msgChanKafka)

	go kafkaWorker(msgChanKafka, usecase)

	// RabbitMQ Consumer
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, usecase)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	fmt.Println("Kafka worker has started")

	for msg := range msgChan {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDTO)
		if err != nil {
			panic(err)
		}

		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Kafka has processed order %s\n", outputDto.ID)
	}
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc usecase.CalculateFinalPrice) {
	fmt.Println("RabbitMQ worker has started")

	for msg := range msgChan {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &OrderInputDTO)
		if err != nil {
			panic(err)
		}

		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		
		fmt.Printf("Rabbitmq has processed order %s\n", outputDto.ID)
	}
}
