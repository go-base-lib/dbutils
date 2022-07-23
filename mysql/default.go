package mysql

import (
	"context"
	dbutils "github.com/byzk-worker/go-db-utils"
	"github.com/jinzhu/gorm"
)

func Db() *gorm.DB {
	return client.Db()
}

func EnableDebug() dbutils.GormCommonInterface {
	return client.EnableDebug()
}

func AutoMigrate(v ...interface{}) dbutils.GormCommonInterface {
	return client.AutoMigrate(v...)
}

func WithContext(callBack func(ctx context.Context) error) error {
	return client.WithContext(callBack)
}
