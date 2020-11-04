package routes
import (
	"gin-api/app/http/controller/dealer"
	"github.com/gin-gonic/gin"
)

func Dealer(router * gin.Engine){
	group := router.Group("/dealer")
	group.GET("/index", dealer.IndexApi)
}