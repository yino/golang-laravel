package redis

import (
	"context"
	"gin-api/config"
	"gin-api/extend/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Client struct {
	Connect *redis.Client
	Ctx     context.Context
}

var Redis *Client

func InitConnectionPool() *Client {
	log.Println("创建 Redis 连接池")
	connect := connection()
	ctx := context.Background()
	Redis = &Client{
		Connect: connect,
		Ctx:     ctx,
	}
	return Redis
}

// 连接数据库
func connection() *redis.Client {
	redisConfig := config.RedisConf()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig["Host"].(string) + ":" + strconv.Itoa(redisConfig["Port"].(int)),
		Password: redisConfig["Password"].(string), // no password set
		DB:       redisConfig["Db"].(int),          // use default DB
		// 连接池容量及闲置连接
		PoolSize:     redisConfig["RedisPoolSize"].(int),
		MinIdleConns: redisConfig["MinIdleConns"].(int),

		// 超时设置
		IdleCheckFrequency: redisConfig["IdleCheckFrequency"].(time.Duration),
		IdleTimeout:        redisConfig["IdleTimeout"].(time.Duration),
		MaxConnAge:         redisConfig["MaxConnAge"].(time.Duration),

		// 命令执行失败重试策略
		MaxRetries:      redisConfig["MaxRetries"].(int),
		MinRetryBackoff: redisConfig["MinRetryBackoff"].(time.Duration),
		MaxRetryBackoff: redisConfig["MaxRetryBackoff"].(time.Duration),
	})
	return rdb
}

// 获取 连接的数据库信息
func Connection() *redis.Client {
	return Redis.Connect
}
func Ctx() context.Context {
	return Redis.Ctx
}
