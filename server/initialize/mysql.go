package initialize

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server-monitor-admin/global"
	"time"
)

//初始化数据库连接
func Mysql() {
	mysql := global.CONFIG.Mysql
	//打开数据库连接
	if db, err := gorm.Open("mysql", mysql.Username+":"+mysql.Password+"@("+mysql.Path+")/"+mysql.Dbname+"?"+mysql.Config); err != nil {
		global.LOG.Error("数据库启动异常", err)
	} else {
		//不默认添加表名后缀s
		db.SingularTable(true)
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		global.DB = db
		global.DB.DB().SetMaxIdleConns(mysql.MaxIdleConns)
		global.DB.DB().SetMaxOpenConns(mysql.MaxOpenConns)
		global.DB.LogMode(mysql.LogMode)
	}

}

//创建钩子
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", time.Now())
	}
}
