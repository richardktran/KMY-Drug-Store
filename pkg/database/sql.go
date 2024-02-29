package database

import (
	"fmt"
	"net/url"

	"github.com/richardktran/KMY-Drug-Store/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSQL() gorm.Dialector {
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

	return mysql.Open(dns)
}
