package model

import (
	"server-monitor-admin/model/base"
)

// 项目
type SysProject struct {
	base.BaseModel
	ProjectName  string `json:"projectName"`  //项目名
	Secret       string `json:"secret"`       // 秘钥
	Description  string `json:"description"`  // 描述
	CreateUserID uint   `json:"createUserId"` //创建人
	DelFlag      int8   `json:"delFlag"`      // 是否删除
}
