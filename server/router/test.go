package router

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor-admin/api/v1"
)

/**
初始化测试方法路由
*/
func InitTestRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("test")
	//其中的变量是单独的作用域
	{
		routerGroup.GET("/test1", v1.Test1)
		routerGroup.GET("/test2", v1.Test2)
	}

}
