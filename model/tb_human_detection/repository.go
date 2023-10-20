package tb_human_detection

import "gorm.io/gorm"

type Repository interface {
	Create(tbHumanDetection TbHumanDetection) (TbHumanDetection, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(tbHumanDetection TbHumanDetection) (TbHumanDetection, error) {
	err := r.db.Create(&tbHumanDetection).Error
	if err != nil {
		return tbHumanDetection, err
	}

	return tbHumanDetection, nil
}
