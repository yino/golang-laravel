package routes

import (
	"github.com/gin-gonic/gin"
	"web-gin-first/app/http/corp/controller"
	_ "web-gin-first/app/http/corp/controller"
)

func Corp(router * gin.Engine) * gin.Engine{
	router.GET("/corp/index", controller.IndexApi)
	return router
}