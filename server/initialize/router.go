package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "server-monitor-admin/docs"
	"server-monitor-admin/global"
	"server-monitor-admin/router"
)

//初始化所有路由
func Routers() *gin.Engine {
	Router := gin.Default()
	//swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//添加统一前缀
	ApiGroup := Router.Group("/api")
	router.InitBaseRouter(ApiGroup)
	router.InitTestRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitGroupRouter(ApiGroup)
	router.InitClientRouter(ApiGroup)
	router.InitItemRouter(ApiGroup)
	global.LOG.Info("router register success")
	return Router
}
