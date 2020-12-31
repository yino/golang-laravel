package main

import (
	"go-laravel/app/exception"
	"go-laravel/app/http/middleware"
	"go-laravel/app/models"
	"go-laravel/config"
	"go-laravel/extend/autoload"
	"go-laravel/extend/redis"
	"go-laravel/routes"

	"github.com/gin-gonic/gin"
)

//go:generate go run main.go
func main() {
	// 自动加载
	autoload.Run()
	// 关闭连接
	defer redis.Connection().Close()
	defer models.Db().Close()
	// 创建新的会话
	router := gin.Default()
	// 跨域等 header问题
	router.Use(middleware.Core())
	// 返回值
	router.Use(middleware.OperationRecords())
	// 接收异常
	router.Use(exception.Recover)
	router = routes.Routers(router)
	router.Run(config.AppPort)
	//defer DB.Close()
}
