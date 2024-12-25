package system

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户表
type SysUser struct {
	gorm.Model // gorm.Model是gorm内置的模型，包含了ID、CreatedAt、UpdatedAt、DeletedAt四个字段
	// 用户UUID
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID" example:"d0123a1b-6682-543f-bdc6-62f333ffc666"`
	// 用户登录名
	Username string `json:"userName" gorm:"comment:用户登录名" example:"admin"`
	// 用户登录密码
	Password string `json:"password"  gorm:"comment:用户登录密码" example:"6ec063004a6f31642261936a212fde3f"`
	// 用户昵称
	NickName string `json:"nickName" gorm:"comment:用户昵称" example:"扣叮侠"`
	// 用户头像
	HeaderImg string `json:"headerImg" gorm:"default:/uploads/head.png;comment:用户头像" example:"https://www.it666.chat/nj_header.jpg"`
	// 用户邮箱
	Email string `json:"email" gorm:"comment:用户邮箱" example:"admin@it666.chat"`
	// 用户手机号
	Phone string `json:"phone" gorm:"comment:用户手机号" example:"18888888888"`
	// 用户openid
	OpenId string `json:"openId" gorm:"comment:用户openid" example:"o6_bmjrPTlm6_2sgVt7hMZOPfL2M"`
}
