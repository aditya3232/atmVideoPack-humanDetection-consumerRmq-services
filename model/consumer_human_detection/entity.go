package consumer_human_detection

// json di struct ini disesuaikan dengan key payload rmq
type RmqConsumerHumanDetection struct {
	TidID                              *int   `json:"tid_id"`
	DateTime                           string `json:"date_time"`
	Person                             string `json:"Person"`
	ConvertedFileCaptureHumanDetection string `json:"converted_file_capture_human_detection"`
}
