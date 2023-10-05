package connection

import (
	"log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

/*
	- setiap koneksi selesai digunakan, agar rapi maka close koneksinya
	- defer close lebih baik digunakan menutup koneksi saat berintaraksi dgn rmq, jgn pas inisialisasi nya
	- defer conn.Close(), defer akan dieksekusi terakhir
*/

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://" + config.CONFIG.RABBIT_USER + ":" + config.CONFIG.RABBIT_PASS + "@" + config.CONFIG.RABBIT_HOST + ":" + config.CONFIG.RABBIT_PORT + "/")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Print("RabbitMQ is connected")

	return conn, nil
}

func RabbitMQ() *amqp.Connection {
	return connection.rabbitmq
}
