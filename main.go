package main

import (
	"gin-api/extend/autoload"
	"gin-api/extend/redis"
	"github.com/gin-gonic/gin"

	"gin-api/app/exception"
	"gin-api/app/models"
	"gin-api/config"
	"gin-api/routes"
)

func main() {
	// 自动加载
	autoload.Run()
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
