package main

import (
	"github.com/jinzhu/gorm"
	"server-monitor-admin/utils"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	v := utils.MD5V([]byte("admin"))
	print(v)
}
