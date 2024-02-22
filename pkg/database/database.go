package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/richardktran/MyBlogBE/pkg/env"
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
		dbHost := env.GET("DB_HOST")
		dbPort := env.GET("DB_PORT")
		dbUser := env.GET("DB_USERNAME")
		dbPass := env.GET("DB_PASSWORD")
		dbName := env.GET("DB_DATABASE")
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
