package database

import (
	"fmt"

	"github.com/richardktran/KMY-Drug-Store/pkg/app"
	"github.com/richardktran/KMY-Drug-Store/pkg/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() gorm.Dialector {
	dbName := env.GET("DB_DATABASE")

	dns := fmt.Sprintf("%s/database/%s.sqlite", app.RootPath(), dbName)

	return sqlite.Open(dns)
}
