package database

import (
	"fmt"

	"github.com/richardktran/MyBlogBE/pkg/app"
	"github.com/richardktran/MyBlogBE/pkg/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() gorm.Dialector {
	dbName := env.GET("DB_DATABASE")
	dns := fmt.Sprintf("%s/database/%s.sqlite", app.RootPath(), dbName)

	return sqlite.Open(dns)
}
