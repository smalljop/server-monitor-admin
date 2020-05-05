package router

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor-admin/api/v1"
	"server-monitor-admin/middleware"
)

//初始化用户路由
func InitUserRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/user").Use(middleware.JwtAuth())
	//其中的变量是单独的作用域
	{
		routerGroup.GET("/info", v1.UserInfo)
	}

}

//初始化项目路由
func InitProjectRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/project").Use(middleware.JwtAuth())
	//其中的变量是单独的作用域
	{
		routerGroup.GET("/list", v1.ListProject)
	}

}
