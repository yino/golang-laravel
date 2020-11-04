package dealer


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func IndexApi(c *gin.Context){
	c.String(http.StatusOK,"dealer !!")
}