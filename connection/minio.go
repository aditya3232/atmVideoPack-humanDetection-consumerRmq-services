package connection

import (
	"log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio() (*minio.Client, error) {
	minioClient, err := minio.New(config.CONFIG.MINIO_HOST+":"+config.CONFIG.MINIO_PORT, &minio.Options{
		Creds:  credentials.NewStaticV4(config.CONFIG.MINIO_ACCESS_KEY, config.CONFIG.MINIO_SECRET_KEY, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Print("Minio is connected")
	return minioClient, nil

}

func Minio() *minio.Client {
	return connection.minio
}
