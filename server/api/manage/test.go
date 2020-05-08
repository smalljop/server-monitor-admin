package manage

import (
	"github.com/gin-gonic/gin"
	"server-monitor-admin/global"
	"server-monitor-admin/global/response"
)

func Test1(c *gin.Context) {
	println(global.VIPER)
	println(global.VIPER.GetString("test.name"))
	response.OkMsg(global.VIPER.GetString("test.name"), c)
}

func Test2(c *gin.Context) {
	c.Get("")
	//var acuser model.AcUser
	//find := global.DB.Find(&acuser)
	//marshal, _ := json.Marshal(&find)
	//println(marshal)
	//response.OkData(find, c)
}
