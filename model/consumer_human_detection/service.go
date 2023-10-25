package consumer_human_detection

type Service interface {
	ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error)
}

type service struct {
	humanDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error) {

	// consume queue
	newRmqConsumerHumanDetection, err := s.humanDetectionRepository.ConsumerQueueHumanDetection()
	if err != nil {
		return newRmqConsumerHumanDetection, err
	}

	return newRmqConsumerHumanDetection, nil

}
