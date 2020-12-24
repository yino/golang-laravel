package routes

import (
	"github.com/gin-gonic/gin"
)


// 路由执行文件
func Routers(route *gin.Engine) *gin.Engine{

	// 加载自定义路由
	Corp(route)
	Dealer(route)

	return route
}