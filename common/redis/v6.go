package redis

import (
	"chatgpt-web-new-go/common/config"
	"github.com/go-redis/redis"
)

func Init() {
	redisConfig := config.Config.Redis
	config.Redis = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // Redis 服务器没有设置密码
		DB:       redisConfig.DB,       // 使用默认数据库
	})

	_, err := config.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
