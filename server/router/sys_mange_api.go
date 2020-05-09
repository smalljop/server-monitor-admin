package router

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor-admin/api/manage"
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
func InitGroupRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/server/group").Use(middleware.JwtAuth())
	//其中的变量是单独的作用域
	{
		routerGroup.GET("/list", v1.ListServerGroup)
		routerGroup.POST("/save", v1.SaveServerGroup)
		routerGroup.POST("/update", v1.UpdateServerGroup)
		routerGroup.POST("/delete", v1.DeleteServerGroup)
	}

}

//初始化项目路由
func InitItemRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/server/item").Use(middleware.JwtAuth())
	//其中的变量是单独的作用域
	{
		routerGroup.GET("/list", v1.ListServerItem)
		routerGroup.POST("/save", v1.SaveServerGroup)
		routerGroup.POST("/update", v1.UpdateServerGroup)
		routerGroup.POST("/delete", v1.DeleteServerGroup)
	}

}
