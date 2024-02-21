package database

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	connection *gorm.DB
}

var dbInstance *database

func GetDB() *gorm.DB {
	if dbInstance == nil {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USERNAME")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_DATABASE")
		dbOptions := url.Values{
			"charset":   {"utf8mb4"},
			"parseTime": {"True"},
			"loc":       {"Local"},
		}
		connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		dns := fmt.Sprintf("%s?%s", connection, dbOptions.Encode())
		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatal("failed to connect database", err)
		}

		dbInstance = &database{connection: db}
	}

	return dbInstance.connection
}

func CloseDB() {
	if db, _ := GetDB().DB(); db != nil {
		err := db.Close()
		if err != nil {
			log.Fatal("failed to close database", err)
		}
	}
}
