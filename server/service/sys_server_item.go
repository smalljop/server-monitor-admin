package service

import (
	"errors"
	"server-monitor-admin/global"
	"server-monitor-admin/model"
	"server-monitor-admin/model/request"
)

// @title    注册机器到服务端
// @description
// @auth    smalljop  2020/4/27 11:27
// @param secret 秘钥
// @param stat 客户端状态信息
// @return

func RegisterServerItem(secret string, stat request.Stat) error {
	group := model.SysServerGroup{}
	//查询秘钥是否正确
	noExists := global.DB.Where("secret=?", secret).First(&group).RecordNotFound()
	if noExists {
		return errors.New("秘钥错误，请注册平台获取正取秘钥")
	}
	item := model.SysServerItem{}
	noItemExists := global.DB.Where("ip_addr=?", stat.IpAddr).First(&item).RecordNotFound()
	if noItemExists {
		item.GroupId = group.Id
		item.IpAddr = stat.IpAddr
		item.HostName = stat.HostName
		global.DB.Create(&item)
	} else {
		global.DB.Update(&item)
	}
	return nil
}
