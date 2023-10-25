package tb_human_detection

import (
	"strconv"
	"time"
)

// entity TbHumanDetection
type TbHumanDetection struct {
	ID                            int       `gorm:"primaryKey" json:"id"`
	CreatedAt                     time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt                     time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	TidID                         *int      `json:"tid_id"`
	DateTime                      string    `json:"date_time"`
	Person                        string    `json:"person"`
	FileNameCaptureHumanDetection string    `json:"file_name_capture_human_detection"`
}

func (m *TbHumanDetection) TableName() string {
	return "tb_human_detection"
}

func (e *TbHumanDetection) RedisKey() string {
	if e.ID == 0 {
		return "tb_human_detection"
	}

	return "tb_human_detection:" + strconv.Itoa(e.ID)
}
