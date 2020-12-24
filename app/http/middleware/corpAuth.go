package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @title Corp login auth
// @desc 登录验证
func CorpLoginAuth() gin.HandlerFunc{
	return func(request *gin.Context) {
		token := request.Query("token")
		fmt.Println(token)
		return
	}
}
