package main

import (
	amqp "github.com/rabbitmq/amqp091-go" 
	"log"
	//"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("Criar a connexao")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitmMQ")
	}

	log.Println("Deu bom filho!", conn)
}