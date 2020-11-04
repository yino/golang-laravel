package routes

import (
	"gin-api/app/http/controller/corp"
	"github.com/gin-gonic/gin"
)

func Corp(router *gin.Engine) {
	group := router.Group("/corp")
	group.GET("/index", corp.IndexApi)
}
