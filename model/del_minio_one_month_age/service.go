package del_minio_one_month_age

import (
	"context"
	"strings"
	"time"

	"github.com/aditya3232/gatewatchApp-services.git/log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/minio/minio-go/v7"
)

type Service interface {
	DelOneMonthOldFiles() error
	DelOneDayOldFiles() error
}

type service struct {
}

// Inisialisasi layanan dengan klien Minio yang sudah ada
func NewService() *service {
	return &service{}
}

func (s *service) DelOneMonthOldFiles() error {
	oneMonthAgo := time.Now().AddDate(0, -1, 0) // Waktu satu bulan yang lalu

	// Membuat daftar objek di dalam bucket
	objectCh := connection.Minio().ListObjects(context.Background(), config.CONFIG.MINIO_BUCKET, minio.ListObjectsOptions{
		Recursive: true,
	})

	for obj := range objectCh {
		if obj.Err != nil {
			return obj.Err
		}

		if obj.LastModified.Before(oneMonthAgo) && obj.Key == "human-detection/"+obj.Key {
			err := connection.Minio().RemoveObject(context.Background(), config.CONFIG.MINIO_BUCKET, obj.Key, minio.RemoveObjectOptions{})
			if err != nil {
				return err
			}
			log.Printf("File %s dihapus\n", obj.Key)
		}
	}
	return nil
}

func (s *service) DelOneDayOldFiles() error {
	// oneDayAgo := time.Now().AddDate(0, 0, -1) // Waktu satu hari yang lalu

	// tenMinuteAgo
	tenMinutesAgo := time.Now().Add(-1 * time.Minute)

	// Membuat daftar objek di dalam bucket
	objectCh := connection.Minio().ListObjects(context.Background(), config.CONFIG.MINIO_BUCKET, minio.ListObjectsOptions{
		Recursive: true,
	})

	for obj := range objectCh {
		if obj.Err != nil {
			return obj.Err
		}

		if obj.LastModified.Before(tenMinutesAgo) && strings.HasPrefix(obj.Key, "human-detection/") {
			err := connection.Minio().RemoveObject(context.Background(), config.CONFIG.MINIO_BUCKET, obj.Key, minio.RemoveObjectOptions{})
			if err != nil {
				return err
			}
			log.Printf("File %s dihapus\n", obj.Key)
		}
	}
	return nil
}
