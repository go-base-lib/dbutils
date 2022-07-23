package sqlite

import (
	dbutils "github.com/byzk-worker/go-db-utils"
)

type Client struct {
	url string
	*dbutils.GormCommon
}

func (c *Client) Init() error {
	return c.GormCommon.Init("sqlite3", c.url)
}
