package manage

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"server-monitor-admin/global"
	"server-monitor-admin/global/response"
	"server-monitor-admin/model"
	"server-monitor-admin/model/request"
	"server-monitor-admin/service"
	"server-monitor-admin/utils"
)

// @Tags Project
// @Summary 项目列表
// @Produce  application/json
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/server/group/list [get]
func ListServerGroup(c *gin.Context) {
	var R request.QueryServerGroup
	_ = c.ShouldBindQuery(&R)
	R.UserId = c.GetInt64(utils.USER_ID_FILED)
	service.ListServerGroup(&R)
	response.OkData(R, c)
}

// @Tags Project
// @Summary 保存项目
// @Produce  application/json
// @Param data body model.SysServerGroup true "对象"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router  /api/server/group/save [get]
func SaveServerGroup(c *gin.Context) {
	var R model.SysServerGroup
	c.ShouldBindJSON(&R)
	userId := c.GetInt64(utils.USER_ID_FILED)
	//生成uuid作为连接秘钥
	v4 := uuid.NewV4()
	R.Secret = v4.String()
	R.CreateUserID = uint(userId)
	global.DB.Create(&R)
	response.Ok(c)
}

// @Tags Project
// @Summary 修改项目
// @Produce  application/json
// @Param data body model.SysServerGroup true "对象"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/server/group/update [post]
func UpdateServerGroup(c *gin.Context) {
	var R model.SysServerGroup
	_ = c.ShouldBindJSON(&R)
	global.DB.Model(&R).Updates(model.SysServerGroup{ProjectName: R.ProjectName, Description: R.Description})
	response.Ok(c)
}

// @Tags Project
// @Summary 删除
// @Produce  application/json
// @Param data body model.SysServerGroup true "对象"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/server/group/delete [post]
func DeleteServerGroup(c *gin.Context) {
	var R model.SysServerGroup
	_ = c.ShouldBindJSON(&R)
	global.DB.Model(&R).Updates(model.SysServerGroup{DelFlag: 1}).Where("id = ?", R.Id)
	response.Ok(c)
}
