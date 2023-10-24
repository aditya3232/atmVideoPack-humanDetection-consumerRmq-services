package del_old_log_from_elastic

type Service interface {
	DelOneMonthOldLogs() error
}

type service struct {
	delOldLogFromElasticRepository Repository
}

func NewService(delOldLogFromElasticRepository Repository) *service {
	return &service{delOldLogFromElasticRepository}
}

func (s *service) DelOneMonthOldLogs() error {
	err := s.delOldLogFromElasticRepository.DelOneMonthOldLogs()
	if err != nil {
		return err
	}

	return nil
}
