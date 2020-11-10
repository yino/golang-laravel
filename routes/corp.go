package routes

import (
	"gin-api/app/http/controller/corp"
	"gin-api/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func Corp(router *gin.Engine) {
	// 创建 corp 分组
	group := router.Group("/corp")
	group.POST("login",corp.LoginController)
	group.Use(middleware.CorpLoginAuth())
	{
		group.GET("/index", corp.IndexController)
	}

	testGroup := router.Group("/corp/test")
	testGroup.GET("/redis", corp.RedisController)


}
