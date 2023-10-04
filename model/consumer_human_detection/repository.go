package consumer_human_detection

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository interface {
	ConsumerQueueHumanDetection() (HumanDetection, error)
	Create(humanDetection HumanDetection) (HumanDetection, error)
}

type repository struct {
	db       *gorm.DB
	rabbitmq *amqp.Connection
}

func NewRepository(db *gorm.DB, rabbitmq *amqp.Connection) *repository {
	return &repository{db, rabbitmq}
}

func (r *repository) ConsumerQueueHumanDetection() (HumanDetection, error) {

	// create channel
	channel, err := r.rabbitmq.Channel()
	if err != nil {
		return HumanDetection{}, err
	}
	defer channel.Close()

	// consume queue
	msgs, err := channel.Consume(
		"HumanDetectionQueue", // name queue
		"",                    // Consumer name (empty for random name)
		true,                  // Auto-acknowledgment (set to true for auto-ack)
		false,                 // Exclusive
		false,                 // No-local
		false,                 // No-wait
		nil,                   // Arguments
	)

	if err != nil {
		return HumanDetection{}, err
	}

	// get message
	for d := range msgs {
		humanDetection := HumanDetection{}
		err := json.Unmarshal(d.Body, &humanDetection)
		if err != nil {
			return HumanDetection{}, err
		}

		// insert Tid, DateTime, Person, File, ConvertedFile from message to db
		_, err = r.Create(
			HumanDetection{
				Tid:                           humanDetection.Tid,
				DateTime:                      humanDetection.DateTime,
				Person:                        humanDetection.Person,
				FileNameCaptureHumanDetection: humanDetection.FileNameCaptureHumanDetection,
			},
		)
		if err != nil {
			return HumanDetection{}, err
		}
	}

	return HumanDetection{}, nil

}

func (r *repository) Create(humanDetection HumanDetection) (HumanDetection, error) {
	err := r.db.Create(&humanDetection).Error
	if err != nil {
		return HumanDetection{}, err
	}

	return humanDetection, nil
}
