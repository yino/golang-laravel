package config

import (
	"reflect"
	"time"
)

var (
	RedisHost     string = "127.0.0.1"
	RedisPassword string = "123456"
	RedisPort     int    = 6379
	RedisDb       int    = 0 // default database
	RedisPrefix	  string = "gin_api_" // 前缀

	RedisPoolSize int = 4  // 连接池最大连接数，默认为4倍CPU数，4* runtime.NumCpu
	MinIdleConns  int = 10 //默认启动连接数，最少连接数

	//闲置连接检查包括
	IdleCheckFrequency time.Duration = 60 * time.Second //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	IdleTimeout        time.Duration = 5 * time.Minute  //闲置超时，默认5分钟，-1表示取消闲置超时检查
	MaxConnAge         time.Duration = 0 * time.Second  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

	// 命令执行失败 重试策略
	MaxRetries      int           = 0                      // 命令执行失败时 最多重试多少次 0不重试
	MinRetryBackoff time.Duration = 8 * time.Millisecond   // 命令重试间隔下限 8毫秒
	MaxRetryBackoff time.Duration = 512 * time.Millisecond // 命令重试间隔上限 512毫秒

)

type redisConfig struct {
	Host     string
	Password string
	Port     int
	Db       int
	RedisPrefix       string

	RedisPoolSize int // 最大连接数
	MinIdleConns  int // 最小连接数

	IdleCheckFrequency time.Duration // 闲置连接检查周期，
	IdleTimeout        time.Duration // 闲置连接超时时间
	MaxConnAge         time.Duration // 连接存活时长

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
}

func RedisConf() map[string]interface{} {
	config := redisConfig{
		Host:     RedisHost,
		Password: RedisPassword,
		Port:     RedisPort,
		Db:       RedisDb,
		RedisPrefix:       RedisPrefix,

		RedisPoolSize: RedisPoolSize,
		MinIdleConns:  MinIdleConns,

		IdleCheckFrequency: IdleCheckFrequency,
		IdleTimeout:        IdleTimeout,
		MaxConnAge:         MaxConnAge,

		MaxRetries:      MaxRetries,
		MinRetryBackoff: MinRetryBackoff,
		MaxRetryBackoff: MaxRetryBackoff,
	}
	elem := reflect.ValueOf(&config).Elem()

	relType := elem.Type()
	data := make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}
