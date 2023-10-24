package del_minio_one_month_age

type Service interface {
	DelOneMonthOldFiles() error
}

type service struct {
	delOneMonthOldFilesRepository Repository
}

func NewService(delOneMonthOldFilesRepository Repository) *service {
	return &service{delOneMonthOldFilesRepository}
}

func (s *service) DelOneMonthOldFiles() error {
	err := s.delOneMonthOldFilesRepository.DelOneMonthOldFiles()
	if err != nil {
		return err
	}

	return nil

}
