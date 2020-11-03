package main

import (
	"fmt"
	"web-gin-first/routes"
	"path/filepath"
)

func main(){
	fmt.Println(filepath.Dir("./routes"))

	router := routes.Router()
	router.Run(":8000")
}