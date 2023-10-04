package main

// func main() {
// 	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
// 	if err != nil {
// 		// return nil, err
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	ch, err := connection.RabbitMQ().Channel()
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	defer ch.Close()

// 	msgs, err := ch.Consume(
// 		"TestQueue",
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	forever := make(chan bool)
// 	go func() {
// 		for d := range msgs {
// 			fmt.Println("Received Message: %s\n", d.Body)
// 		}
// 	}()

// 	fmt.Println("Successfully connected to our RabbitMQ instance")
// 	fmt.Println(" [*] - waiting for messages")
// 	<-forever

// }
