package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-monitor-admin/global"
	"server-monitor-admin/global/response"
	"server-monitor-admin/utils"
	"strconv"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("X-Token")
		if len(auth) == 0 {
			context.Abort()
			response.FailCodeMsg(http.StatusUnauthorized, "请重新登录", context)
		}
		// 校验token
		c, err := parseToken(auth)
		if c != nil {
			userId, _ := strconv.ParseInt(c.Id, 10, 64)
			context.Set(utils.USER_ID_FILED, userId)
		}
		if err != nil {
			context.Abort()
			response.FailCodeMsg(http.StatusUnauthorized, "token 过期"+err.Error(), context)
		}
		context.Next()
	}
}

//解析token
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.CONFIG.Jwt.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
