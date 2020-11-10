package main

import (
	"fmt"
	"gin-api/extend/redis"
	"github.com/gin-gonic/gin"

	"gin-api/app"
	"gin-api/app/exception"
	"gin-api/app/models"
	"gin-api/app/validate"
	"gin-api/config"
	"gin-api/extend/log"
	"gin-api/routes"
)

func main() {

	log.SetOutPutDir(app.GetLogDir())
	// log 模块
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutPutDir(app.GetLogDir())
	// 检测是否支持中文扩展 （用于验证器返回中文）
	if err := validate.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// 创建MySql连接池
	models.InitDbConnection()
	// 创建redis 连接池
	redis.InitConnectionPool()

	// 关闭连接
	defer redis.Connection().Close()
	defer models.Db().Close()
	// 创建新的会话
	router := gin.Default()
	router.Use(exception.Recover)
	router = routes.Routers(router)
	router.Run(config.AppPort)
	//defer DB.Close()
}
