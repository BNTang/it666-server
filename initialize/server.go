package initialize

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"it666.com/global"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 1.初始化路由
	Router := Routers()
	// 2.拿到配置文件中的端口号
	address := fmt.Sprintf(":%d", global.IT666_CONFIG.System.Port)
	// 3.根据端口号和路由启动服务
	s := initServer(address, Router)
	global.IT666_LOG.Info("服务启动成功, 端口号:", zap.String("address", address))
	global.IT666_LOG.Error(s.ListenAndServe().Error())
}

// 初始化服务
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
