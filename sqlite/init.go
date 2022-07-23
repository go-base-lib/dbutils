package sqlite

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

func InitWithOptions(localFile string, fn dbutils.GetContextFn) error {
	return Init(fmt.Sprintf("file:%s?auto_vacuum=1", localFile), fn)
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
