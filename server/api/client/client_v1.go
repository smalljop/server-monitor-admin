package client

import (
	"github.com/gin-gonic/gin"
	//"server-monitor-admin/global"
	"server-monitor-admin/global/response"
	"server-monitor-admin/model/request"
	"server-monitor-admin/service"
	"server-monitor-admin/utils"
)

// @Tags Client
// @Summary 加入项目
// @Produce  application/json
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/client/join/serverGroup [get]
func JoinServerGroup(c *gin.Context) {
	secret := c.PostForm("secret")
	stat := request.Stat{}
	c.ShouldBindJSON(&stat)
	err := service.RegisterServerItem(secret, stat)
	response.IsOK(err, c)
}

// @Tags Client
// @Summary 同步状态信息
// @Produce  application/json
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/client/join/project [post]
func SyncClientStatInfo(c *gin.Context) {
	//维护到内存中 加快查询速度
	stat := request.Stat{}
	c.ShouldBindJSON(&stat)
	utils.ClientStatCache[stat.IpAddr] = stat
	//marshal, _ := json.Marshal(&stat)
	response.Ok(c)
}
