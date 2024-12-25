package system

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"it666.chat/global"
	"it666.chat/model/system"
	"it666.chat/utils"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	// 判断用户名是否注册
	var user system.SysUser
	if !errors.Is(global.IT666_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 用户昵称
	if u.NickName == "" {
		u.NickName = u.Username
	}
	// 密码加密
	u.Password = utils.MD5V([]byte(u.Password + global.IT666_CONFIG.System.Salt))
	// 设置uuid
	u.UUID = uuid.NewV4()
	// 插入用户到数据库
	err = global.IT666_DB.Create(&u).Error
	return u, err
}
