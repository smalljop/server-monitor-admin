package service

import (
	"server-monitor-admin/global"
	"server-monitor-admin/model"
	"server-monitor-admin/model/request"
	"server-monitor-admin/utils"
)

func ListServerGroup(page *request.QueryServerGroup) {
	var projects []model.SysServerGroup
	db := global.DB.Limit(page.Limit).Offset(page.GetOffset()).Order("create_time desc")
	if page.UserId != utils.SUPER_ADMIN_USER_ID {
		db = db.Where("create_user_id=?", page.UserId)
	}
	if len(page.ProjectName) != 0 {
		db = db.Where("project_name like ?", "%"+page.ProjectName+"%")
	}
	db.Where("del_flag=0").Find(&projects).Count(&page.Total)
	page.List = projects
}
