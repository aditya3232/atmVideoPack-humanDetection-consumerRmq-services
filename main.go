package main

import (
	"fmt"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/model/consumer_human_detection"
	"github.com/gin-gonic/gin"
)

func init() {
	forever := make(chan bool)
	go func() {
		consumerHumanDetectionRepository := consumer_human_detection.NewRepository(connection.DatabaseGatewatch(), connection.RabbitMQ())
		consumerHumanDetectionService := consumer_human_detection.NewService(consumerHumanDetectionRepository)

		_, err := consumerHumanDetectionService.ConsumerQueueHumanDetection()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(" [*] - waiting for messages")
	<-forever

}

func main() {
	// panic recovery
	defer helper.RecoverPanic()

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))

}

// request
// handler = yg menangani input, service, repository, entity
// service = fungsi bisnis prosesnya
// input = handle inputan
// repository = query
// entity = struct dari table

// helper = func bantuan yg sering dipakai
// formatter = untuk memformat response
// auth = isinya generate & validate token

// penulisan interface repository
// - (create, find, update, delete) bisa diimbuhi create mau crud tabel apa atau data apa

// penulisan interface service
// - (create, find/get, update, delete) bisa diimbuhi create mau crud tabel apa atau data apa
// - kata kerja lain ex: (RegisterUser, Login, IsEmailAvailable)

// handler bisa sama dengan service, lebih baik sama biar g bingung

// authorization
// ambil nilai header authorization: Bearer tokentoken
// dari header authorization, kita ambil nilai tokennya saja
// kita validasi token
// kita ambil user_id
// kita ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user, agar infomasi user dapat diakses kemana aja, context adalah sebuah tempat untuk menyimpan suatu nilai, yg akhirnya dapat di get di tempat lain
