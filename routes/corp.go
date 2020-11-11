package routes

import (
	"github.com/gin-gonic/gin"
	"go-laravel/app/http/controller/corp"
	"go-laravel/app/http/middleware"
)

func Corp(router *gin.Engine) {
	// 创建 corp 分组
	group := router.Group("/corp")
	group.POST("login", corp.LoginController)
	group.Use(middleware.CorpLoginAuth())
	{
		group.GET("/index", corp.IndexController)
	}

	testGroup := router.Group("/corp/test")
	testGroup.GET("/redis", corp.RedisController)

}
