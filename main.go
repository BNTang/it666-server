package main

import (
	"it666.com/global"
	"it666.com/initialize"
)

/*
后端项目目录结构
|--it666-server

	    |---api: api接口
		|---config: 配置文件结构体
		|---initialize: 初始化(路由、数据库、配置文件等)
		|---middleware: 中间件
		|---model: 模型(存放数据库模型代码)
		|---router: 路由(存放路由代码)
		|---service: 服务层(存放业务逻辑代码)
		|---utils: 工具类(存放工具类代码)
		|---global: 全局变量(存放全局变量代码)
		|---logs: 日志文件(存放各种日志文件)
		|---main.go: 入口文件
		|---go.mod: go mod配置文件
		|---config(.yaml/.json): 配置文件
*/
func main() {
	// 初始化viper
	global.IT666_VIPER = initialize.Viper("yaml")
	// 初始化zap日志库
	logDir := global.IT666_CONFIG.Zap.Director
	isConsole := global.IT666_CONFIG.Zap.LogInConsole
	format := global.IT666_CONFIG.Zap.Format
	global.IT666_LOG = initialize.Zap(logDir, isConsole, format)
	// 初始化mysql数据库
	global.IT666_DB = initialize.Gorm()
	// 注册数据库表
	if global.IT666_DB != nil {
		initialize.RegisterTables(global.IT666_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.IT666_DB.DB()
		defer db.Close()
	}
	// 启动服务
	initialize.RunServer()
}
