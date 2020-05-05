package model

import (
	"server-monitor-admin/model/base"
)

type SysUser struct {
	base.BaseModel
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickName"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
}
