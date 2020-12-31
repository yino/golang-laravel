package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func OperationRecords() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 响应记录
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// 开始时间
		startTime := time.Now().Unix()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now().Unix()
		// 执行时间
		latencyTime := endTime - startTime
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		params := make(map[string]interface{})
		params["ip"] = clientIP
		params["method"] = reqMethod
		params["latency"] = int(latencyTime)
		params["path"] = reqUri
		params["status"] = statusCode
		params["startTime"] = int(startTime)
		params["endTime"] = int(endTime)
		params["agent"] = c.Request.UserAgent()
		params["resp"] = bodyLogWriter.body.String()
		corpInfo, err := c.Get("CorpInfo")
		if err != false {
			CorpInfo := corpInfo.(map[string]interface{})
			params["userId"] = CorpInfo["id"].(int)
		}
		if reqMethod == "GET" {
			params["body"] = reqUri
		} else if reqMethod == "POST" {
			form, _ := json.Marshal(c.Request.PostForm)
			params["body"] = string(form)
		}
		//fmt.Println(params)
		//operationRecordsService := new(service.OperationRecordsService)
		//operationRecordsService.Add(params)
	}
}
