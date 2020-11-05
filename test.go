package main

import (
	"fmt"
	"gin-api/app/models"
	_ "os"
	_ "reflect"
)

func main() {
	var corpAdmin []models.CorpAdminModel
	models.Db().Find(&corpAdmin)
	for _, item := range corpAdmin {
		fmt.Println(item.Account)
	}
	//fmt.Println(corpAdmin)
}
