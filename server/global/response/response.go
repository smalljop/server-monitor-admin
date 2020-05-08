package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	//状态码
	Code int `json:"code"`
	//响应数据
	Data interface{} `json:"data"`
	//提示消息
	Msg string `json:"msg"`
}

const (
	SUCCESS = 200 //成功
	ERROR   = 500 //失败
)

//返回结果
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

//根据err判断是否成功
func IsOK(err error, c *gin.Context) {
	if err != nil {
		Ok(c)
	} else {
		FailMsg(err.Error(), c)
	}
}

//成功
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}
func OkData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetail(msg string, data interface{}, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func FailCodeMsg(code int, msg string, c *gin.Context) {
	Result(code, map[string]interface{}{}, msg, c)
}
