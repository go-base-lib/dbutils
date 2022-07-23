package mysql

import (
	"fmt"
	dbutils "github.com/byzk-worker/go-db-utils"
)

var client *Client

func Init(url string, getContextFn dbutils.GetContextFn) error {
	c := New(url, getContextFn)
	if err := c.Init(); err != nil {
		return err
	}

	client = c
	return nil
}

func InitWithOptions(username, password, address, dbName string, fn dbutils.GetContextFn) error {
	return Init(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName), fn)
}

func New(url string, getContextFn dbutils.GetContextFn) *Client {
	c := &Client{
		url: url,
		GormCommon: &dbutils.GormCommon{
			CtxFn: getContextFn,
		},
	}
	return c
}
