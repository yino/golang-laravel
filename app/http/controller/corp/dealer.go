package corp

import (
	"fmt"
	"go-laravel/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(c *gin.Context) {
	var corpAdmin models.CorpAdminModel
	var account string
	account = "admin"
	go models.Db().Where("account = ?", account).First(&corpAdmin)
	fmt.Println(corpAdmin.Account)
	c.String(http.StatusOK, "It works")
}



