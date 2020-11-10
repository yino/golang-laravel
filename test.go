package main

import (
	"gin-api/extend/log"
	"gin-api/extend/redis"
	"time"
)

func main() {
	// redis 连接池
	redis.InitConnectionPool()
	redisConnect := redis.Connection()
	err := redisConnect.Set("abc","zba",2 * time.Second).Err()

	if err !=nil {
		log.Println("redis写入失败")
	}
	log.Println("写入成功")

	log.Println("测试")
}
