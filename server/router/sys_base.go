package router

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor-admin/api/v1"
)

//初始化用户路由
func InitBaseRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/")
	//其中的变量是单独的作用域
	{
		routerGroup.POST("user/register", v1.Register)
		routerGroup.POST("user/login", v1.Login)
	}

}
