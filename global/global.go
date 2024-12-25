package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"it666.com/config"
)

var (
	// IT666_VIPER 是一个全局的Viper实例
	IT666_VIPER *viper.Viper
	// IT666_CONFIG 是一个全局的配置实例
	IT666_CONFIG config.Server
	// IT666_LOG 是一个全局的zap日志实例
	IT666_LOG *zap.Logger
	// IT666_DB 是一个全局的MySQL数据库实例
	IT666_DB *gorm.DB
	// IT666_REDIS 是一个全局的Redis数据库实例
	IT666_REDIS *redis.Client
)
