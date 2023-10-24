package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/connection"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	logrus "github.com/sirupsen/logrus"
)

var New = logrus.New()

// func init() {
// 	log := New

// 	os.Chdir("../atmVideoPack-services/log")
// 	file, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Warnf("error opening file: %v", err)
// 	}

// 	mw := io.MultiWriter(os.Stdout, file)
// 	log.SetOutput(mw)

// 	log.Formatter = &logrus.JSONFormatter{
// 		TimestampFormat: "15:04:05 02-01-2006",
// 		FieldMap: logrus.FieldMap{
// 			logrus.FieldKeyTime:  "timestamp",
// 			logrus.FieldKeyLevel: "level",
// 			logrus.FieldKeyMsg:   "message",
// 		},
// 		DisableHTMLEscape: false,
// 	}
// }

// sendLogToElasticsearch sends the log to Elasticsearch
func sendLogToElasticsearch(level logrus.Level, args ...interface{}) {
	// Create a buffer for the JSON payload
	var buf bytes.Buffer
	// Create a map for the JSON payload
	var data map[string]interface{}
	var err error

	// Create the JSON payload from the logrus fields
	data = make(map[string]interface{})
	// data["fields"] = logrus.Fields{}
	data["level"] = level.String()
	data["message"] = args[0]
	data["timestamp"] = time.Now().Format("15:04:05 02-01-2006")

	// Add the args to the JSON payload
	if len(args) > 1 {
		data["args"] = args[1:]
	}

	// Serialize the data to JSON
	if err = json.NewEncoder(&buf).Encode(data); err != nil {
		panic(err)
	}

	// Generate a random UUID as the DocumentID
	randomID := uuid.New().String()

	// Set up the request object directly from the Elasticsearch library
	req := esapi.IndexRequest{
		Index:      "human_detection_publisher_log",
		DocumentID: randomID, // make it random
		Body:       &buf,
		Refresh:    "true",
	}

	// Perform the request with the client
	res, err := req.Do(context.Background(), connection.ElasticSearch())
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	defer res.Body.Close()
}

func Info(args ...interface{}) {
	New.Info(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Infof(format string, args ...interface{}) {
	New.Infof(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Error(args ...interface{}) {
	New.Error(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Errorf(format string, args ...interface{}) {
	New.Errorf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Fatal(args ...interface{}) {
	New.Fatal(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Fatalf(format string, args ...interface{}) {
	New.Fatalf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Panic(args ...interface{}) {
	New.Panic(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Panicf(format string, args ...interface{}) {
	New.Panicf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	New.Warn(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Warnf(format string, args ...interface{}) {
	New.Warnf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Debug(args ...interface{}) {
	New.Debug(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Debugf(format string, args ...interface{}) {
	New.Debugf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Trace(args ...interface{}) {
	New.Trace(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Tracef(format string, args ...interface{}) {
	New.Tracef(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Print(args ...interface{}) {
	New.Print(args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Printf(format string, args ...interface{}) {
	New.Printf(format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Log(level logrus.Level, args ...interface{}) {
	New.Log(level, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	New.Logf(level, format, args...)
	sendLogToElasticsearch(logrus.InfoLevel, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return New.WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return New.WithField(key, value)
}

func WithError(err error) *logrus.Entry {
	return New.WithError(err)
}
