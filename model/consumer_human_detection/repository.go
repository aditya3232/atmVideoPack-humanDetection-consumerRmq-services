package consumer_human_detection

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/helper"
	libraryMinio "github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/library/minio"
	log_function "github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/model/add_human_detection_to_elastic"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository interface {
	ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error)
}

type repository struct {
	db            *gorm.DB
	rabbitmq      *amqp.Connection
	elasticsearch *esv7.Client
}

func NewRepository(db *gorm.DB, rabbitmq *amqp.Connection, elasticsearch *esv7.Client) *repository {
	return &repository{db, rabbitmq, elasticsearch}
}

func (r *repository) ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error) {
	var rmqConsumerHumanDetection RmqConsumerHumanDetection

	// create channel
	ch, err := r.rabbitmq.Channel()
	if err != nil {
		return rmqConsumerHumanDetection, err
	}
	defer ch.Close()

	// consume queue
	msgs, err := ch.Consume(
		"HumanDetectionQueue", // name queue
		"",                    // Consumer name (empty for random name)
		true,                  // Auto-acknowledgment (set to true for auto-ack)
		false,                 // Exclusive
		false,                 // No-local
		false,                 // No-wait
		nil,                   // Arguments
	)

	if err != nil {
		return rmqConsumerHumanDetection, err
	}

	// get message
	for d := range msgs {
		newHumanDetection := rmqConsumerHumanDetection
		err := json.Unmarshal(d.Body, &newHumanDetection)
		if err != nil {
			return rmqConsumerHumanDetection, err
		}

		// konversi humanDetection.FileNameCaptureHumanDetection string ke bytes
		bytesConvertedFile, err := base64.StdEncoding.DecodeString(newHumanDetection.ConvertedFileCaptureHumanDetection)
		if err != nil {
			return rmqConsumerHumanDetection, err
		}

		// Mengunggah gambar ke MinIO
		objectName := "human-detection/" + helper.DateTimeToStringWithStrip(time.Now()) + ".jpg"
		FileNameCaptureHumanDetection := helper.DateTimeToStringWithStrip(time.Now()) + ".jpg"

		key, err := libraryMinio.UploadFileFromPutObject(config.CONFIG.MINIO_BUCKET, objectName, bytesConvertedFile)
		if err != nil {
			log_function.Error(fmt.Sprintf("Gambar gagal diunggah ke MinIO dengan nama objek: %s\n", key.Key))
			return rmqConsumerHumanDetection, err
		}

		// add data newHumanDtection to elasticsearch with CreateElasticHumanDetection
		repoElastic := add_human_detection_to_elastic.NewRepository(r.elasticsearch)
		resultElastic, err := repoElastic.CreateElasticHumanDetection(
			add_human_detection_to_elastic.ElasticHumanDetection{
				ID:                            helper.DateTimeToStringWithStrip(time.Now()),
				TidID:                         newHumanDetection.TidID,
				DateTime:                      newHumanDetection.DateTime,
				Person:                        newHumanDetection.Person,
				FileNameCaptureHumanDetection: FileNameCaptureHumanDetection,
			},
		)
		if err != nil {
			return rmqConsumerHumanDetection, err
		}
		// log result elastic
		log_function.Info(fmt.Sprintf("Result elastic: %v\n", resultElastic))

		// create data tb_human_detection
		// repo := tb_human_detection.NewRepository(r.db)
		// _, err = repo.Create(
		// 	tb_human_detection.TbHumanDetection{
		// 		TidID:                         newHumanDetection.TidID,
		// 		DateTime:                      newHumanDetection.DateTime,
		// 		Person:                        newHumanDetection.Person,
		// 		FileNameCaptureHumanDetection: FileNameCaptureHumanDetection,
		// 	},
		// )
		// if err != nil {
		// 	return rmqConsumerHumanDetection, err
		// }

	}

	return rmqConsumerHumanDetection, nil

}
