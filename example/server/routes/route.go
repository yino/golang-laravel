package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-laravel/config"
)


// 路由执行文件
func Routers(route *gin.Engine) *gin.Engine{
	//route.GET("/status")
	route.StaticFS("/storage",http.Dir(config.StaticDir))
	route.GET("/status", func(context *gin.Context) {
		context.String(200," ok!")
	})
	// 加载自定义路由
	Corp(route)
	return route
}