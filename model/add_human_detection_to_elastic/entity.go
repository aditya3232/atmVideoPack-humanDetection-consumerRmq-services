package add_human_detection_to_elastic

// ini entity data yg akan dikirim ke elastic

type ElasticHumanDetection struct {
	ID                            string `json:"id"`
	Tid                           string `json:"tid"`
	DateTime                      string `json:"date_time"`
	Person                        string `json:"person"`
	FileNameCaptureHumanDetection string `json:"file_name_capture_human_detection"`
}
