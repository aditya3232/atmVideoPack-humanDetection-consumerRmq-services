package del_minio_one_month_age

import (
	"context"
	"time"

	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/config"
	log_function "github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/log"
	"github.com/minio/minio-go/v7"
)

type Repository interface {
	DelOneMonthOldFiles() error
}

type repository struct {
	minio *minio.Client
}

func NewRepository(minio *minio.Client) *repository {
	return &repository{minio}
}

func (r *repository) DelOneMonthOldFiles() error {
	oneMonthAgo := time.Now().AddDate(0, -1, 0) // Waktu satu bulan yang lalu

	// Membuat daftar objek di dalam bucket
	objectCh := r.minio.ListObjects(context.Background(), config.CONFIG.MINIO_BUCKET, minio.ListObjectsOptions{
		Recursive: true,
	})

	for obj := range objectCh {
		if obj.Err != nil {
			return obj.Err
		}

		if obj.LastModified.Before(oneMonthAgo) && obj.Key == "human-detection/"+obj.Key {
			err := r.minio.RemoveObject(context.Background(), config.CONFIG.MINIO_BUCKET, obj.Key, minio.RemoveObjectOptions{})
			if err != nil {
				return err
			}
			log_function.Printf("File %s dihapus\n", obj.Key)
		}
	}
	return nil
}
