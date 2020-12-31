package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"go-laravel/app/exception"
	"go-laravel/app/models"
	"go-laravel/config"
	"go-laravel/extend/log"
	"go-laravel/extend/redis"
	"time"
)

func CorpLoginService(account string, password string) {
	redisCorpTokenPre := "corp_token:"
	var corpAdmin models.CorpAdminModel
	result := models.Db().Where("account = ?", account).First(&corpAdmin)
	// 查询结果为空
	if result.Error != nil {
		exception.BaseException(50001, "账号不存在")
	}
	// 禁用
	if corpAdmin.Status == models.CorpAdminStatusClose {
		exception.BaseException(50001, "账号被禁用")
	}
	// 验证密码
	authPassword := corpAdmin.Salt + password
	h := md5.New()
	h.Write([]byte(authPassword))
	authPassword = hex.EncodeToString(h.Sum(nil))
	if corpAdmin.Password != authPassword {
		exception.BaseException(50000, "密码错误")
	}

	token := corpAdmin.Token
	// 根据 token 去redis中查询 是否 有token
	_,err := AuthTokenService(token)

	data := make(map[string]string)
	data["username"] = corpAdmin.Name
	data["account"] = corpAdmin.Account
	// 重新生成token
	if err != nil {
		// 生成token
		baseService := new(BaseService)
		token, err := baseService.CreateUserJwtToken(string(corpAdmin.ID), "CORP_ADMIN")
		if err != nil {
			exception.BaseException(50000, "token生成失败")
		}

		// user info
		userInfo := make(map[string]interface{})
		userInfo["username"] = corpAdmin.Name
		userInfo["account"] = corpAdmin.Account
		userInfo["id"] = corpAdmin.ID
		jsonUserInfo, _ := json.Marshal(userInfo)
		// 将token 存储到数据库中
		redis.Connection().Set(redis.Ctx(), redisCorpTokenPre+token, string(jsonUserInfo), config.LoginTokenExpire)
		corpAdmin.Token = token
	}

	// 更新最后登录时间
	corpAdmin.LastLoginTime = int(time.Now().Unix())
	result = models.Db().Save(&corpAdmin)
	data["token"] = token
	// 生成token、session
	exception.SuccessResponse(data)
}

func AuthTokenService(token string) (map[string]interface{}, error) {

	log.Println("----------------------AuthTokenService--------------------")
	redisCorpTokenPre := "corp_token:"
	// 根据 token 去redis中查询 是否 有token
	log.Write("Auth Token:", token)
	result, err := redis.Connection().Get(redis.Ctx(), redisCorpTokenPre+token).Result()
	log.Println("----------------------Auth Token Result：", result, err, "--------------------")

	res := make(map[string]interface{})
	if err != nil {
		return res, errors.New("token过期")
	} else {
		json.Unmarshal([]byte(result), &res)
		return res , err
	}
}
