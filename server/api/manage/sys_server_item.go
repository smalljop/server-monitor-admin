package manage

import (
	"github.com/gin-gonic/gin"
	"server-monitor-admin/global/response"
	"server-monitor-admin/utils"
)

// @Tags Client
// @Summary  查询服务器客户端项
// @Produce  application/json
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/client/join/project [post]
func ListServerItem(c *gin.Context) {
	caches := utils.ClientStatCache
	list := make([]interface{}, 0)
	for _, v := range caches {
		list = append(list, v)
	}
	response.OkData(list, c)
}
