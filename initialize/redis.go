package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"it666.com/global"
)

func Redis() *redis.Client {
	redisCfg := global.IT666_CONFIG.Redis
	// 打开redis数据库
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.IT666_LOG.Error("Redis连接失败, err:", zap.Error(err))
		return nil
	} else {
		global.IT666_LOG.Info("Redis连接成功 response:", zap.String("pong", pong))
		return client
	}
}
