package main

import (
	"fmt"
	kafka2 "github.com/PortoCode/codelivery/simulator/application/kafka"
	"github.com/PortoCode/codelivery/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	// spawn a new thread (asynchronous) to consume
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		// spawn a new thread (asynchronous) to produce
		go kafka2.Produce(msg)
	}
}