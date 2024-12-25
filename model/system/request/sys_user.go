package request

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 用户注册请求参数
type UserRegister struct {
	// 用户登录名
	Username string `form:"userName"  json:"userName"  binding:"required,min=5,max=20" error_info:"用户名必须是5-20位的字符"` // 用户名必须是5-20位的字符
	// 用户登录密码
	Password string `form:"password" json:"password" binding:"required,verifyPwd" error_info:"密码应包含数字、大/小写字母、特殊字符中的3种, 且至少8个字符"` // 密码必须是8-20位的字符
	// 用户昵称
	NickName string `form:"nickName" json:"nickName" binding:"omitempty,min=2,max=20" error_info:"昵称必须是2-20位的字符"` // 昵称必须是2-20位的字符
	// 用户头像
	HeaderImg string `form:"headerImg" json:"headerImg" binding:"omitempty,url" error_info:"头像必须是完整URL"` // 头像必须是完整URL
	// 用户邮箱
	Email string `form:"email" json:"email" binding:"omitempty,email" error_info:"邮箱格式不正确"` // 邮箱格式不正确
	// 用户手机号
	Phone string `form:"phone" json:"phone" binding:"omitempty,verifyPhone" error_info:"手机号格式不正确"` // 手机号格式不正确
	// 用户openid
	OpenId string `form:"openId" json:"openId" binding:"omitempty" error_info:"openid不能为空"` // openid不能为空
}

// 用户密码格式验证
func VerifyPwd(f validator.FieldLevel) bool {
	// 拿到密码值
	val := f.Field().String()
	// 验证密码长度
	if len(val) < 8 || len(val) > 20 {
		return false
	}
	// 验证密码是否包含特殊字符
	pwdPattern := `^[0-9a-zA-Z\W]{8,20}$`
	reg, err := regexp.Compile(pwdPattern)
	if err != nil {
		return false
	}
	// 匹配密码
	match := reg.MatchString(val)
	if !match {
		return false
	}
	// 判断是否包含数字、大小写字母、特殊字符
	var cnt int = 0
	patternList := []string{
		`[0-9]+`,
		`[a-z]+`,
		`[A-Z]+`,
		`[\W]+`,
	}
	for _, pattern := range patternList {
		match, _ = regexp.MatchString(pattern, val)
		if match {
			cnt++
		}
	}
	fmt.Println(cnt)
	// 满足3中以上即可通过验证
	return cnt >= 3
}

// 用户手机号格式验证
func VerifyPhone(f validator.FieldLevel) bool {
	// 拿到密码值
	val := f.Field().String()
	// 验证密码是否包含特殊字符
	pwdPattern := `^(13\d|14[5-9]|15[0-3,5-9]|16[56]|17[0-8]|18\d|19[89])\d{8}$`
	reg, err := regexp.Compile(pwdPattern)
	if err != nil {
		return false
	}
	// 匹配密码
	match := reg.MatchString(val)
	return match
}
