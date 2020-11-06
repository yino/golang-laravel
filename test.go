package main

import (
	"fmt"
)

type name struct {
	name string
	age int
}

func one(name string, age int) int {
	return age
}


func main()  {
	//
	//var arr [10]int
	//var str string
	//var a int
	//var b bool
	//var c float64
	// list dict array map array
	var tMap map[string]int
	tMap["a"] = 1
	var testName name
	testName.name = "sun"
	testName.age = 20
	fmt.Println(testName)
}