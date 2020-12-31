package corp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-laravel/app/exception"
	"go-laravel/app/service"
	"go-laravel/app/validate"
	"go-laravel/extend/log"
)

// 登录控制器
func LoginController(request *gin.Context) {
	// 验证器
	var params validate.CorpLogin
	log.Println("--------------AuthLogin---------------")
	if err := request.ShouldBindJSON(&params); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		log.Println("login：", err)
		validate.ErrResp(errs)
		// validator.ValidationErrors类型错误则进行翻译
	}
	service.CorpLoginService(params.Account, params.Password)
}

func GetUserInfoController(request *gin.Context) {
	result, _ := request.Get("CorpInfo")
	exception.SuccessResponse(result)
}
