package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-laravel/config"
)

// cmd 钩子
type CmdHook struct {
}

//NewLogHook 新建 redis 日志记录 hook
func NewCmdHook() *CmdHook {
	return &CmdHook{}
}

// 命令执行前
func (hook *CmdHook) BeforeProcess(c context.Context, cmd redis.Cmder) (context.Context, error) {
	args := cmd.Args()
	// 重新拼接 key
	if len(args) > 1 {
		redisConfig := config.RedisConf()
		args[1] = redisConfig["RedisPrefix"].(string) + args[1].(string)
		cmd = redis.NewCmd(c, args)
		log.Printf("【redis】%v", args)
	}

	if cmd.Err() != nil {
		log.Println(c, cmd.String(), "cmd 执行错误", cmd.Err())
		return nil, cmd.Err()
	}
	return c, nil
}

func (hook *CmdHook) AfterProcess(c context.Context, cmd redis.Cmder) error {
	fmt.Println(cmd, c)
	return nil
}

func (hook *CmdHook) BeforeProcessPipeline(c context.Context, cmds []redis.Cmder) (context.Context, error) {
	return c, nil
}
func (hook *CmdHook) AfterProcessPipeline(c context.Context, cmds []redis.Cmder) error {
	return nil
}

func Boost(c *redis.Client) *redis.Client {
	c.AddHook(NewCmdHook())
	return c
}
