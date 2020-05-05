package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server-monitor-admin/global"
	"server-monitor-admin/global/response"
	"server-monitor-admin/model"
	"server-monitor-admin/model/request"
	resp "server-monitor-admin/model/response"
	"server-monitor-admin/service"
	"server-monitor-admin/utils"
	"strconv"
	"time"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.RegisterStruct true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/user/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := model.SysUser{NickName: R.NickName, Email: R.Email, Password: R.Password}
	_, err := service.Register(user)
	if err != nil {
		response.FailMsg(fmt.Sprintf("%v", err), c)
	} else {
		response.OkMsg("注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.LoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	var R request.LoginStruct
	_ = c.ShouldBindJSON(&R)
	user := model.SysUser{UserName: R.UserName, Password: R.Password}
	user, err := service.Login(user)
	if err != nil {
		response.FailMsg("用户名或者密码错误", c)
	} else {
		generateJwtToken(user, c)
	}
}

// @Tags Base
// @Summary 获取用户信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/user/info [get]
func UserInfo(c *gin.Context) {
	userId := c.GetInt64(utils.USER_ID_FILED)
	user, err := service.GetUserInfo(userId)
	if err != nil {
		response.FailMsg("获取用户信息失败", c)
	} else {
		var info = make(map[string]interface{})
		info["name"] = user.NickName
		info["avatar"] = user.Avatar
		info["roles"] = "admin"
		response.OkData(info, c)
	}
}

//生成token
func generateJwtToken(user model.SysUser, c *gin.Context) {
	// 省略代码
	expiresTime := time.Now().Unix() + global.CONFIG.Jwt.ExpireTime
	claims := jwt.StandardClaims{
		Audience:  user.UserName,                           // 受众
		ExpiresAt: expiresTime,                             // 失效时间
		Id:        strconv.FormatUint(uint64(user.Id), 10), // 编号
		IssuedAt:  time.Now().Unix(),                       // 签发时间
		Issuer:    global.CONFIG.Jwt.Issuer,                // 签发人
		NotBefore: time.Now().Unix(),                       // 生效时间
		Subject:   "login",                                 // 主题
	}
	var jwtSecret = []byte(global.CONFIG.Jwt.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		response.OkData(resp.LoginResult{
			Token:    token,
			UserName: user.UserName,
			Avatar:   user.Avatar,
		}, c)
	} else {
		response.OkMsg("登录失败", c)
	}

}
