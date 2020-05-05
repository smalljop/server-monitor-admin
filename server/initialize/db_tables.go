package initialize

import (
	"server-monitor-admin/global"
	"server-monitor-admin/model"
)

//注册所有表
func DbTables() {
	db := global.DB
	db.AutoMigrate(&model.SysUser{})
	db.AutoMigrate(&model.SysProject{})
	global.LOG.Debug("register table success")
}
