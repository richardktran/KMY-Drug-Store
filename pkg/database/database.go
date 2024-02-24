package database

import (
	"log"

	"github.com/richardktran/MyBlogBE/pkg/env"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	connection *gorm.DB
}

var dbInstance *database

func GetDB() *gorm.DB {
	if dbInstance == nil {
		db, err := gorm.Open(selectConnectionDB(), &gorm.Config{
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

func selectConnectionDB() gorm.Dialector {
	dbConnection := env.GET("DB_CONNECTION")

	switch dbConnection {
	case "mysql":
		return NewSQL()
	case "sqlite":
		return NewSQLite()
	default:
		return NewSQL()
	}
}
