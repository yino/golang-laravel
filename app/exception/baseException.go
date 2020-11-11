package exception

import (
	"fmt"
	"go-laravel/config"
	"go-laravel/extend/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type exception struct {
	Code int
	Msg  string
	Data interface{}
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("----------recover-------")
			// 打印错误堆栈信息
			// 如果要用到上述结构体 必须 断言一下类型 坑！！！
			log.Printf("panic: %v\n", r.(exception).Code)
			//fmt.Println(reflect.TypeOf(r))
			// 显示错误log
			if config.AppDebug == true {
				debug.PrintStack()
			}
			// 封装通用json返回
			c.JSON(http.StatusOK, gin.H{
				"code": r.(exception).Code,
				"msg":  r.(exception).Msg,
				"data": r.(exception).Data,
			})
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func BaseException(code int, msg string) {

	var err exception

	err.Code = code
	err.Msg = msg
	err.Data = nil
	panic(err)
}

func SuccessResponse(data interface{}) {
	var success exception
	success.Code = 200
	success.Msg = "success"
	success.Data = data
	panic(success)
}
