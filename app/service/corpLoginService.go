package service

import (
	"fmt"
	"gin-api/app/exception"
	"gin-api/app/models"
)

func CorpLoginService(account string, password string){
	fmt.Println(account)
	var corpAdmin models.CorpAdminModel
	result := models.Db().Where("account = ?", account).First(&corpAdmin)

	// 查询结果为空
	if result.Error !=nil {
		exception.BaseException(50001,"账号不存在")
	}

	fmt.Println(account)
	fmt.Println(corpAdmin)

}