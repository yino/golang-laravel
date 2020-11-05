package main

import (
	"fmt"
	"gin-api/app/exception"
	"gin-api/app/validate"
	"gin-api/config"
	"gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// 检测是否支持中文扩展 （用于验证器返回中文）
	if err := validate.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 创建新的会话
	router := gin.Default()
	router.Use(exception.Recover)
	router = routes.Routers(router)
	router.Run(config.AppPort)
}
