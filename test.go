package main

import (
	"gin-api/app"
	"gin-api/extend/log"
)

func main() {

	log.SetOutPutDir(app.GetLogDir())
	log.Println("测试")
}
