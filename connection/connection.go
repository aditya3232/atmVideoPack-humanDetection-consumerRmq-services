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

type Connection struct {
	db       *gorm.DB
	redis    *redis.Client
	es       *elasticsearch.Client
	minio    *minio.Client
	rabbitmq *amqp.Connection
}

var (
	debug      int = config.CONFIG.DEBUG
	connection Connection
	initOnce   sync.Once
)

// untuk matikan koneksi ke database
// - dari init nya
// - dan dari repository nya
// - dan untk elastic di log nya
func init() {
	initOnce.Do(func() {
		db, err := connectDatabaseMysql()
		if err != nil {
			// log.Panic(err)
			panic(err)
		}
		minio, err := ConnectMinio()
		if err != nil {
			panic(err)
		}
		rabbitmq, err := ConnectRabbitMQ()
		if err != nil {
			panic(err)
		}
		// redis, err := ConnectRedis()
		// if err != nil {
		// 	panic(err)
		// }
		// es, err := ConnectElastic()
		// if err != nil {
		// 	panic(err)
		// }

		connection = Connection{
			db:       db,
			minio:    minio,
			rabbitmq: rabbitmq,
			// redis: redis,
			// es: es,
		}
	})
}

func Close() {
	if connection.db != nil {
		sqlDB, _ := connection.db.DB()
		sqlDB.Close()
		connection.db = nil
	}
	if connection.minio != nil {
		connection.minio = nil
	}
	if connection.rabbitmq != nil {
		connection.rabbitmq.Close()
		connection.rabbitmq = nil
	}

	// if connection.redis != nil {
	// 	connection.redis.Close()
	// 	connection.redis = nil
	// }

	// if connection.es != nil {
	// 	connection.es.Close()
	// 	connection.es = nil
	// }
}
