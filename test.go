package main

import (
	"reflect"
	"time"
)

func main() {
	// redis 连接池
	day := time.Now().Month()
	println(reflect.TypeOf(day))
}
