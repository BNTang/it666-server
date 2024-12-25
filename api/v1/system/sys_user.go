package system

import (
	"github.com/gin-gonic/gin"
	"it666.com/model/common/response"
	systemModel "it666.com/model/system"
	systemReq "it666.com/model/system/request"
	"it666.com/utils"
)

type UserApi struct{}

// 注册账号
func (b *UserApi) Register(c *gin.Context) {
	// 接收参数并验证
	var userRegister systemReq.UserRegister
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		msg := utils.ProcessErr(userRegister, err)
		response.FailWithMessage(msg, c)
		return
	}
	// 调用service层注册账号
	user := systemModel.SysUser{
		Username:  userRegister.Username,
		Password:  userRegister.Password,
		NickName:  userRegister.NickName,
		HeaderImg: userRegister.HeaderImg,
		Email:     userRegister.Email,
		Phone:     userRegister.Phone,
		OpenId:    userRegister.OpenId,
	}
	userReturn, err := userService.Register(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.SuccessWithDetailed(userReturn, "注册成功", c)
}

// 登录账号
func (b *UserApi) Login(c *gin.Context) {
	response.SuccessWithMessage("登录成功", c)
}

// 删除账号
func (b *UserApi) DeleteUser(c *gin.Context) {
}

// 更新账号
func (b *UserApi) UpdateUserInfo(c *gin.Context) {
}

// 获取账号信息
func (b *UserApi) GetUserInfo(c *gin.Context) {
}
