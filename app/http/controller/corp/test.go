package corp

import (
	"go-laravel/extend/log"
	"go-laravel/extend/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RedisController(c *gin.Context)  {
	key := "test"
	err := redis.Connection().Set(redis.Ctx(),key,"asdasd",10*time.Second).Err()
	if err != nil{
		log.Panic("redis 写入错误")
	}

	result,err := redis.Connection().Get(redis.Ctx(),key).Result()
	if err != nil{
		log.Panic("redis get fail")
	}
	c.String(http.StatusOK,"SUCCESS, Result："+result)
}