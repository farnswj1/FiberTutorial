package cache

import (
	"app/utils"

	"github.com/gofiber/storage/redis/v3"
)

var Redis *redis.Storage

func loadRedis() {
	if Redis != nil {
		return
	}

	Redis = redis.New(redis.Config{
		URL: utils.Env["REDIS_URL"],
	})
}

func init() {
	loadRedis()
}
