package config

import (
	"log"

	"github.com/spf13/viper"
)

var CONFIG = load()

type envConfigs struct {
	DEBUG int `mapstructure:"DEBUG"`

	DB_1_HOST    string `mapstructure:"DB_1_HOST"`
	DB_1_PORT    string `mapstructure:"DB_1_PORT"`
	DB_1_USER    string `mapstructure:"DB_1_USER"`
	DB_1_PASS    string `mapstructure:"DB_1_PASS"`
	DB_1_NAME    string `mapstructure:"DB_1_NAME"`
	DB_1_CHARSET string `mapstructure:"DB_1_CHARSET"`
	DB_1_LOC     string `mapstructure:"DB_1_LOC"`

	APP_PORT string `mapstructure:"APP_PORT"`
	API_KEY  string `mapstructure:"API_KEY"`
	APP_HOST string `mapstructure:"APP_HOST"`

	// elasticsearch
	ES_HOST string `mapstructure:"ES_HOST"`
	ES_PORT string `mapstructure:"ES_PORT"`

	// redis
	REDIS_HOST string `mapstructure:"REDIS_HOST"`
	REDIS_PORT string `mapstructure:"REDIS_PORT"`
	REDIS_PASS string `mapstructure:"REDIS_PASS"`

	// minio
	MINIO_HOST       string `mapstructure:"MINIO_HOST"`
	MINIO_PORT       string `mapstructure:"MINIO_PORT"`
	MINIO_ACCESS_KEY string `mapstructure:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY string `mapstructure:"MINIO_SECRET_KEY"`
	MINIO_BUCKET     string `mapstructure:"MINIO_BUCKET"`

	// rabbitmq
	RABBIT_HOST string `mapstructure:"RABBIT_HOST"`
	RABBIT_PORT string `mapstructure:"RABBIT_PORT"`
	RABBIT_USER string `mapstructure:"RABBIT_USER"`
	RABBIT_PASS string `mapstructure:"RABBIT_PASS"`
}

func load() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("../../atmVideoPack-publisherRmq-services/config")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
