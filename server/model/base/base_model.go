package base

type BaseModel struct {
	Id         uint      `gorm:"primary_key"json:"id"`
	CreateTime LocalTime `json:"createTime"`
	UpdateTime LocalTime `json:"updateTime"`
}
