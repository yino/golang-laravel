package autoload

import (
	"fmt"
	"go-laravel/app"
	"go-laravel/app/models"
	"go-laravel/app/validate"
	"go-laravel/extend/log"
	"go-laravel/extend/redis"
)

// 加载log
func loadLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutPutDir(app.GetLogDir())
}

// 自动加载
func Run() {
	// 加载log
	loadLog()
	// 中文验证器
	verifyZn()
	// mysql pool
	mysqlInitPool()
	// redis pool
	redisInitPool()
}

// 加载中文验证器

func verifyZn() {
	// 检测是否支持中文扩展 （用于验证器返回中文）
	if err := validate.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
}

func mysqlInitPool()  {
	// 创建MySql连接池
	models.InitDbConnection()
}
func redisInitPool()  {
	// 创建redis 连接池
	redis.InitConnectionPool()
}
