package main

// func main() {
// 	// tes rabbitmq
// 	ch, err := connection.RabbitMQ().Channel()
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"TestQueue",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	fmt.Println(q)

// 	err = ch.Publish(
// 		"",
// 		"TestQueue",
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte("Hello World!"),
// 		},
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	fmt.Println("Successfully published message to queue")

// 	// panic recovery
// 	defer helper.RecoverPanic()

// 	router := gin.Default()
// 	if config.CONFIG.DEBUG == 0 {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	// router.Use(gin.Recovery())

// 	routes.Initialize(router)
// 	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))

// }
