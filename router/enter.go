package router

import (
	"it666.com/router/bbs"
	"it666.com/router/system"
)

type RouterGroup struct {
	// 初始化系统路由
	System system.RouterGroup
	// 初始化论坛路由
	Bbs bbs.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
