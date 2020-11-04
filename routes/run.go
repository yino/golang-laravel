package routes

import (
	"github.com/gin-gonic/gin"
)
func Routers() *gin.Engine{
	route := gin.Default()

	Corp(route)
	Dealer(route)

	return route
}