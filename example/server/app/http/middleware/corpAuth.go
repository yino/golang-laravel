package middleware

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"go-laravel/app/exception"
	"go-laravel/app/service"
	"strings"
)

// @title Corp login auth
// @desc 登录验证
func CorpLoginAuth() gin.HandlerFunc {
	return func(request *gin.Context) {
		token := request.GetHeader("Authorization")
		if len(token) < 1 {
			exception.BaseException(-1, "请登录")
		}
		token = strings.Split(token, "Bearer ")[1]
		data, err := service.AuthTokenService(token)
		if err != nil {
			exception.BaseException(-1, "token 失效")
		} else {
			data["id"] = int(data["id"].(float64))
			request.Set("CorpInfo", data)
		}
		return
	}
}
