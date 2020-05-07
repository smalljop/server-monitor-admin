package request

import "server-monitor-admin/model/base"

type QueryProject struct {
	base.BasePage
	UserId      int64  `json:"userId"`
	ProjectName string `form:"projectName" json:"projectName"`
}
