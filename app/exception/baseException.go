package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"gin-api/config"
	"runtime/debug"
)

type exception struct {

	Code int
	Msg string
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("----------recover-------")
			// 打印错误堆栈信息
			// 如果要用到上述结构体 必须 断言一下类型 坑！！！
			log.Printf("panic: %v\n", r.(exception).Code)
			fmt.Println(reflect.TypeOf(r))
			// 显示错误log
			if config.AppDebug == true {
				debug.PrintStack()
			}
			// 封装通用json返回
			c.JSON(http.StatusOK, gin.H{
				"code": r.(exception).Code,
				"msg": r.(exception).Msg,
				"data": nil,
			})
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func BaseException(code int, msg string)  {

	var err exception

	err.Code = code
	err.Msg = msg
	panic(err)
}