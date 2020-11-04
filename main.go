package main

import (
	"gin-api/routes"
	"gin-api/config"

)

func main(){
	router := routes.Routers()
	router.Run(config.AppPort)
}