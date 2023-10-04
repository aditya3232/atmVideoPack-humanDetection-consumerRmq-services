package tes_rmq

import (
	"fmt"

	"github.com/aditya3232/gatewatchApp-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	Queue()
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) Queue() {
	// do something
	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to queue")
}
