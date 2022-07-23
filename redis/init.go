package redis

import (
	dbutils "github.com/byzk-worker/go-db-utils"
	redisV8 "github.com/go-redis/redis/v8"
)

var client *Client

func Init(options *redisV8.Options, getContextFn dbutils.GetContextFn) {
	client = New(options, getContextFn)
}

func InitWithOptions(options *redisV8.Options) {
	client = New(options, dbutils.DefaultGetContextFn)
}

func New(options *redisV8.Options, getContextFn dbutils.GetContextFn) *Client {
	c := redisV8.NewClient(options)
	return &Client{
		context: getContextFn,
		c:       c,
	}
}
