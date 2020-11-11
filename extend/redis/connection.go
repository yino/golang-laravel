package redis

import (
	"context"
	"go-laravel/config"
	"go-laravel/extend/log"
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

	//可连接性检测
	_, err := connect.Ping(ctx).Result()
	if err != nil {
		log.Panic(err)
	}


	poolstats := connect.PoolStats()
	log.Printf("redis 总连接数=%d,空闲连接数=%d,已经移除的连接数=%d\n",
		poolstats.TotalConns,
		poolstats.IdleConns,
		poolstats.StaleConns)

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
