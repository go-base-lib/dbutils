package dbutils

import (
	"context"
	"github.com/jinzhu/gorm"
)

type GetContextFn func() (context.Context, error)

var DefaultGetContextFn GetContextFn = func() (context.Context, error) {
	return context.Background(), nil
}

type GormCommonInterface interface {
	Db() *gorm.DB
	EnableDebug() GormCommonInterface
	AutoMigrate(v ...interface{}) GormCommonInterface
	WithContext(callBack func(ctx context.Context) error) error
}

type GormCommon struct {
	db    *gorm.DB
	CtxFn GetContextFn
}

func (c *GormCommon) Init(dialect, url string) error {
	var err error

	c.db, err = gorm.Open(dialect, url)
	if err != nil {
		return err
	}
	return nil
}

func (c *GormCommon) Db() *gorm.DB {
	return c.db
}

func (c *GormCommon) EnableDebug() GormCommonInterface {
	c.db = c.db.Debug()
	return c
}

func (c *GormCommon) AutoMigrate(v ...interface{}) GormCommonInterface {
	c.db.AutoMigrate(v...)
	return c
}

func (c *GormCommon) WithContext(callBack func(ctx context.Context) error) error {
	ctx, err := c.CtxFn()
	if err != nil {
		return err
	}
	return callBack(ctx)
}
