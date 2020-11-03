package main

import (
	"fmt"
	"io/ioutil"
	"log"
	_ "os"
	_ "reflect"
	"strings"
)

func main(){
	fileInfoList, err := ioutil.ReadDir("./routes")
	if err != nil {
		log.Fatal(err)
	}
	// 包含了 run.go 文件
	if len(fileInfoList) <1{
		panic("请编写路由文件")
	}
	for i := range fileInfoList {
		// 文件列表
		fileName := fileInfoList[i].Name()
		//reflect.TypeOf(fileInfoList[i].Name())
		fileArr := strings.Split(fileName,".go")
		if len(fileArr) <1{
			continue
		}
		fmt.Println(fileArr[0])
		//fmt.Println(app.Capitalize(fileArr[0].(string)))
	}
}
