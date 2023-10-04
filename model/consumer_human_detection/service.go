package consumer_human_detection

type Service interface {
	ConsumerQueueHumanDetection() (HumanDetection, error)
}

type service struct {
	humanDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueHumanDetection() (HumanDetection, error) {

	// consume queue
	humanDetection, err := s.humanDetectionRepository.ConsumerQueueHumanDetection()
	if err != nil {
		return HumanDetection{}, err
	}

	return humanDetection, nil

}
