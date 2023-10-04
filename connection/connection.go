package connection

import (
	"sync"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/elastic/go-elasticsearch"
	"github.com/minio/minio-go/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Database struct {
	db       *gorm.DB
	redis    *redis.Client
	es       *elasticsearch.Client
	minio    *minio.Client
	rabbitmq *amqp.Connection
}

var (
	debug    int = config.CONFIG.DEBUG
	database Database
	initOnce sync.Once
)

// untuk matikan koneksi ke database
// - dari init nya
// - dan dari repository nya
// - dan untk elastic di log nya
func init() {
	initOnce.Do(func() {
		db, err := connectDatabaseGatewatch()
		if err != nil {
			// log.Panic(err)
			panic(err)
		}
		minio, err := ConnectMinioGatewatch()
		if err != nil {
			panic(err)
		}
		rabbitmq, err := ConnectRabbitMQ()
		if err != nil {
			panic(err)
		}
		// redis, err := ConnectRedisGatewatch()
		// if err != nil {
		// 	panic(err)
		// }
		// es, err := ConnectElastic()
		// if err != nil {
		// 	panic(err)
		// }

		database = Database{
			db:       db,
			minio:    minio,
			rabbitmq: rabbitmq,
			// redis: redis,
			// es: es,
		}
	})
}

func Close() {
	if database.db != nil {
		sqlDB, _ := database.db.DB()
		sqlDB.Close()
		database.db = nil
	}
	if database.minio != nil {
		database.minio = nil
	}
	if database.rabbitmq != nil {
		database.rabbitmq.Close()
		database.rabbitmq = nil
	}

	// if database.redis != nil {
	// 	database.redis.Close()
	// 	database.redis = nil
	// }

	// if database.es != nil {
	// 	database.es.Close()
	// 	database.es = nil
	// }
}
