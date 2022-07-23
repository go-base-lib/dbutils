package mysql

import (
	dbutils "github.com/byzk-worker/go-db-utils"
)

type Client struct {
	url string
	*dbutils.GormCommon
}

func (c *Client) Init() error {
	return c.GormCommon.Init("mysql", c.url)
}
