package tes_rmq

type Service interface {
	TesQueueRmq()
}

type service struct {
	tesRmqRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) TesQueueRmq() {
	s.tesRmqRepository.Queue()
}
