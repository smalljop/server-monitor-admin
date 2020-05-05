package service

import (
	"server-monitor-admin/global"
	"server-monitor-admin/model"
	"server-monitor-admin/model/base"
)

func ListProject(page *base.BasePage) {
	var projects []model.SysProject
	global.DB.Limit(page.Limit).Offset(page.GetOffset()).Order("create_time desc").Find(&projects).Count(&page.Total)
	page.List = projects
}
