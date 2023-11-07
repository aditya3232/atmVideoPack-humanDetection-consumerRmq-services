package consumer_human_detection

// json di struct ini disesuaikan dengan key payload rmq
type RmqConsumerHumanDetection struct {
	Tid                                string `json:"tid"`
	DateTime                           string `json:"date_time"`
	Person                             string `json:"Person"`
	ConvertedFileCaptureHumanDetection string `json:"converted_file_capture_human_detection"`
}
