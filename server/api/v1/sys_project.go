package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor-admin/global/response"
	"server-monitor-admin/model/base"
	"server-monitor-admin/service"
)

// @Tags Project
// @Summary 项目列表
// @Produce  application/json
// @Param data body base.BasePage true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/project/list [get]
func ListProject(c *gin.Context) {
	var R base.BasePage
	_ = c.ShouldBindQuery(&R)
	service.ListProject(&R)
	response.OkData(R, c)
}
