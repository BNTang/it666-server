package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	v1 "it666.chat/api/v1"
	systemReq "it666.chat/model/system/request"
)

type UserRouter struct{}

func (s *UserRouter) InitPublicUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.UserApi

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("verifyPwd", systemReq.VerifyPwd)
		v.RegisterValidation("verifyPhone", systemReq.VerifyPhone)
	}
	// 用户注册账号
	userRouter.POST("register", baseApi.Register)
	// 用户登录
	userRouter.POST("login", baseApi.Login)
}

func (s *UserRouter) InitPrivateUserRouter(Router *gin.RouterGroup) {

	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	userRouter.PUT("update", baseApi.UpdateUserInfo) // 更新用户信息
	userRouter.GET("info", baseApi.GetUserInfo)      // 获取用户信息
	userRouter.DELETE("delete", baseApi.DeleteUser)  // 删除用户
}
