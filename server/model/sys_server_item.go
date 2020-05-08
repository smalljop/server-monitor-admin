package model

import "server-monitor-admin/model/base"

// 客户端服务器像
type SysServerItem struct {
	base.BaseModel
	GroupId  uint   `json:"groupId"`  // 服务器组Id
	HostName string `json:"hostName"` // hostName
	IpAddr   string `json:"ipAddr"`   // ip地址
}
