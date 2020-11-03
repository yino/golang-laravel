package routes

import (
	"github.com/gin-gonic/gin"
)
func Routers() *gin.Engine{
	route := gin.Default()
	route = Corp(route)
}