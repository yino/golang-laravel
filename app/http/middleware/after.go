package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// http response
// @middleware
func after() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() {
			return
		}
		params := c.Keys
		if len(params) == 0 {
			return
		}
		c.JSON(http.StatusOK, params)
	}
}