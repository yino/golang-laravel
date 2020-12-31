package corp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-laravel/extend/redis"
)

func TestController(request *gin.Context)  {
	fmt.Println(redis.Connection().Get(redis.Ctx(),"test"))
}