package connection

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

/*
	- setiap koneksi selesai digunakan, agar rapi maka close koneksinya
	- defer close lebih baik digunakan menutup koneksi saat berintaraksi dgn rmq, jgn pas inisialisasi nya
	- defer conn.Close(), defer akan dieksekusi terakhir
*/

func ConnectRabbitMQ() (*amqp.Connection, error) {
	// conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.CONFIG.RMQ_USER, config.CONFIG.RMQ_PASS, config.CONFIG.RMQ_HOST, config.CONFIG.RMQ_PORT))
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully connected to our RabbitMQ instance")

	return conn, nil
}

func RabbitMQ() *amqp.Connection {
	return database.rabbitmq
}
