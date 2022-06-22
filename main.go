package main

import  (
	"log"
	
	    kafkaApp "github.com/rafaelalmeida/codeedu/imersaofsfc2-simulator/application/kafka"
		ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
		"fmt"
	    "github.com/rafaelalmeida/codeedu/imersaofsfc2-simulator/infra/kafka"
		"github.com/joho/godotenv"
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
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApp.Produce(msg)
	}
}
