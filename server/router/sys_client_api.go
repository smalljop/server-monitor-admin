package router

import (
	"github.com/gin-gonic/gin"
	"server-monitor-admin/api/client"
)

//初始化用户路由
func InitClientRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/client/v1")
	//其中的变量是单独的作用域
	{
		routerGroup.POST("/sync/stat", client.SyncClientStatInfo)
	}

}
