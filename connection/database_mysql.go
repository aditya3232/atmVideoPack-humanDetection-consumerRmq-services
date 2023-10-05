package connection

import (
	"fmt"
	"log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectDatabaseMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", config.CONFIG.DB_1_USER, config.CONFIG.DB_1_PASS, config.CONFIG.DB_1_HOST, config.CONFIG.DB_1_PORT, config.CONFIG.DB_1_NAME, config.CONFIG.DB_1_CHARSET, config.CONFIG.DB_1_LOC)

	logMode := logger.Silent
	if debug == 1 {
		logMode = logger.Info
		// fmt.Println("Database connection string: ", dsn)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Print("Database is connected")
	return db, nil
}

func DatabaseMysql() *gorm.DB {
	return connection.db
}
