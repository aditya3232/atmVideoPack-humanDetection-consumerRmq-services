package main

import (
	"fmt"

	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/connection"
	_ "github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/cron"
	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/helper"
	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/model/consumer_human_detection"
	"github.com/gin-gonic/gin"
)

func main() {
	defer helper.RecoverPanic()

	forever := make(chan bool)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine

		consumerHumanDetectionRepository := consumer_human_detection.NewRepository(connection.DatabaseMysql(), connection.RabbitMQ(), connection.ElasticSearch())
		consumerHumanDetectionService := consumer_human_detection.NewService(consumerHumanDetectionRepository)

		_, err := consumerHumanDetectionService.ConsumerQueueHumanDetection()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(" [*] - waiting for messages")
	<-forever

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))

	// Penutupan koneksi setelah aplikasi selesai berjalan
	defer connection.Close()

}
