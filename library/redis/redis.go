package redis

import (
	"e.coding.net/handnote/handnote/library/setting"
	"github.com/go-redis/redis"
)

// RedisClient redis 连接对象.
var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Addr,
		Password: setting.Redis.Password,
		DB:       setting.Redis.DB,
	})
}