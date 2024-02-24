package database

import (
	"fmt"
	"net/url"

	"github.com/richardktran/MyBlogBE/pkg/app"
	"github.com/richardktran/MyBlogBE/pkg/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() gorm.Dialector {
	dbName := env.GET("DB_DATABASE")
	dbOptions := url.Values{
		"charset":   {"utf8mb4"},
		"parseTime": {"True"},
		"loc":       {"Local"},
	}

	dns := fmt.Sprintf("%s/database/%s.sqlite?%s", app.RootPath(), dbName, dbOptions.Encode())

	return sqlite.Open(dns)
}
