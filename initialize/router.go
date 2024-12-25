package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"it666.chat/global"
	"it666.chat/middleware"
	"it666.chat/router"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()
	// 为用户默认头像提供静态资源服务
	Router.StaticFS(global.IT666_CONFIG.System.Path, http.Dir(global.IT666_CONFIG.System.Path))
	// 跨域规则
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System

	// 无需鉴权就可以访问的路由
	PublicGroup := Router.Group("")
	systemRouter.InitPublicUserRouter(PublicGroup)

	// 需要鉴权才可以访问的路由
	PrivateGroup := Router.Group("")
	systemRouter.InitPrivateUserRouter(PrivateGroup)

	return Router
}
