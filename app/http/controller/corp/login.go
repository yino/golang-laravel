package corp

import (
	_ "fmt"
	"gin-api/app/service"
	"gin-api/app/validate"
	"gin-api/extend/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 登录控制器
func LoginController(request *gin.Context) {
	// 验证器
	account := request.PostForm("account")
	password := request.PostForm("password")
	var corpLoginValidate validate.CorpLogin
	log.Println("AuthLogin")
	if err := request.ShouldBind(&corpLoginValidate); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		request.JSON(http.StatusOK, validate.ErrResp(errs))
		return
	}
	service.CorpLoginService(account, password)
}
