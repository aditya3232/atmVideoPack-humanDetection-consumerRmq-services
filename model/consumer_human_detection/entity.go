package consumer_human_detection

// json di struct ini disesuaikan dengan key payload rmq
type HumanDetection struct {
	Tid                           string `json:"Tid"`
	DateTime                      string `json:"DateTime"`
	Person                        string `json:"Person"`
	FileNameCaptureHumanDetection string `json:"ConvertedFile"`
}

// table name
func (m *HumanDetection) TableName() string {
	return "tb_human_detection"
}
