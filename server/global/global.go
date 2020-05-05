package global

import (
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"server-monitor-admin/config"
)

//全局通用对象
var (
	CONFIG *config.Config    //配置对象
	LOG    *oplogging.Logger //日志
	VIPER  *viper.Viper      //全局viper对象 方便读取配置
	DB     *gorm.DB          //数据库连接对象
)
