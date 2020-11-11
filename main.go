package main

import (
	"github.com/gin-gonic/gin"
	"go-laravel/extend/autoload"
	"go-laravel/extend/redis"

	"go-laravel/app/exception"
	"go-laravel/app/models"
	"go-laravel/config"
	"go-laravel/routes"
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
