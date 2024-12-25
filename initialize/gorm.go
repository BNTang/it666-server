package initialize

import (
	"os"
	"time"

	"it666.chat/global"
	"it666.chat/model/system"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
* 1.什么是gorm
* gorm是一个Go语言的ORM框架，它是一个对象关系映射，是一个对象关系映射器，
* 它可以将数据库中的表映射为结构体，将结构体映射为表中的记录
* 大白话就是不用写SQL语句，就可以操作数据库, 用对象的方式操作数据库
* https://gorm.io/zh_CN/
 */
func Gorm() *gorm.DB {
	// 从全局配置中获取mysql配置
	m := global.IT666_CONFIG.Mysql
	// 如果没有配置mysql数据库，直接返回nil
	if m.Dbname == "" {
		return nil
	}
	// 初始化mysql配置
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	// 打开mysql数据库
	if db, err := gorm.Open(mysql.New(mysqlConfig), getConfig()); err != nil {
		global.IT666_LOG.Error("MySQL启动异常", zap.Error(err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns) // 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(m.MaxOpenConns) // 设置打开数据库连接的最大数量
		sqlDB.SetConnMaxLifetime(time.Hour)   // 设置连接可复用的最大时间
		global.IT666_LOG.Info("MySQL启动成功")
		return db
	}
}

// 注册数据库表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 所有表创建在这里
		system.SysUser{},
	)
	if err != nil {
		global.IT666_LOG.Error("创建数据库表失败", zap.Error(err))
		os.Exit(0)
	}
	global.IT666_LOG.Info("创建数据库表成功")
}
