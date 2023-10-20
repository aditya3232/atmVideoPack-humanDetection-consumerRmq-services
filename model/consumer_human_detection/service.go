package consumer_human_detection

type Service interface {
	ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueHumanDetection() (RmqConsumerHumanDetection, error) {

	// consume queue
	newRmqConsumerHumanDetection, err := s.statusMcDetectionRepository.ConsumerQueueHumanDetection()
	if err != nil {
		return newRmqConsumerHumanDetection, err
	}

	return newRmqConsumerHumanDetection, nil

}
