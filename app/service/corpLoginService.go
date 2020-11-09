package service

import (
	"crypto/md5"
	"encoding/hex"
	"gin-api/app/exception"
	"gin-api/app/models"
	"gin-api/extend/log"
	"time"
)

func CorpLoginService(account string, password string){
	var corpAdmin models.CorpAdminModel
	result := models.Db().Where("account = ?", account).First(&corpAdmin)

	// 查询结果为空
	if result.Error !=nil {
		exception.BaseException(50001,"账号不存在")
	}

	// 禁用
	if corpAdmin.Status == models.CorpAdminStatusClose{
		exception.BaseException(50001,"账号被禁用")
	}

	// 验证密码
	authPassword := password+corpAdmin.Salt
	h := md5.New()
	h.Write([]byte(authPassword))
	authPassword = hex.EncodeToString(h.Sum(nil))
	if corpAdmin.Password != authPassword {
		exception.BaseException(50000,"密码错误")
	}

	// 更新最后登录时间
	corpAdmin.LastLoginTime = int(time.Now().Unix())
	result = models.Db().Save(&corpAdmin)

	data := make(map[string]string)
	data["username"] = corpAdmin.Username
	data["account"] = corpAdmin.Account
	// 生成token、session
	log.Println(data)
	exception.SuccessResponse(data)
}