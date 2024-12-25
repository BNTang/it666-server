package system

import "it666.chat/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
