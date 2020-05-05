package main

import (
	_ "server-monitor-admin/config"
	"server-monitor-admin/core"
	"server-monitor-admin/global"
	"server-monitor-admin/initialize"
)

func main() {
	//初始化mysql
	initialize.Mysql()
	//注册所有表
	initialize.DbTables()
	//程序结束关闭数据库连接
	defer global.DB.Close()
	//启动服务
	core.RunServer()
}
