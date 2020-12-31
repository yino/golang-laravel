package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-laravel/config"
	"go-laravel/extend/log"
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
			log.Printf("panic: %v\n", r)
			// 显示错误log
			if config.AppDebug == true {
				debug.PrintStack()
			}
			response, ok := r.(exception)
			if ok {
				c.JSON(http.StatusOK, gin.H{
					"code": response.Code,
					"msg":  response.Msg,
					"data": response.Data,
				})
			} else {
				// 封装通用json返回
				c.JSON(http.StatusOK, gin.H{
					"code": 40000,
					"msg":  "服务器异常",
					"data": "",
				})
			}
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
func response(c *gin.Context, r exception) {
	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"data": r.Data,
	})
}

func BaseException(code int, msg string, ) {

	var err exception
	err.Code = code
	err.Msg = msg
	err.Data = nil
	panic(err)
}

func ValidateResponse(data interface{}) {
	var err exception
	err.Code = 600000
	err.Msg = "ERROR,参数不全；" + data.(string)
	panic(err)
}

func SuccessResponse(data interface{}) {
	var success exception
	success.Code = 200
	success.Msg = "SUCCESS"
	success.Data = data
	panic(success)
}
