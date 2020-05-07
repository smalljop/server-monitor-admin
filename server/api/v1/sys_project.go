package v1

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
// @Router /api/project/list [get]
func ListProject(c *gin.Context) {
	var R request.QueryProject
	_ = c.ShouldBindQuery(&R)
	R.UserId = c.GetInt64(utils.USER_ID_FILED)
	service.ListProject(&R)
	response.OkData(R, c)
}

// @Tags Project
// @Summary 保存项目
// @Produce  application/json
// @Param data body model.SysProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/project/save [get]
func SaveProject(c *gin.Context) {
	var R model.SysProject
	err := c.ShouldBindJSON(&R)
	global.LOG.Error(err)
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
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/project/update [get]
func UpdateProject(c *gin.Context) {
	var R model.SysProject
	_ = c.ShouldBindJSON(&R)
	global.DB.Model(&R).Updates(model.SysProject{ProjectName: R.ProjectName, Description: R.Description})
	response.Ok(c)
}

// @Tags Project
// @Summary 修改项目
// @Produce  application/json
// @Param data body request.QueryProject true "分页数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/project/delete [get]
func DeleteProject(c *gin.Context) {
	var R request.IdStruct
	_ = c.ShouldBindJSON(&R)
	global.DB.Model(&R).Updates(model.SysProject{DelFlag: 1})
	response.Ok(c)
}
