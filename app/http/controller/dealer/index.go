package dealer


import (
	"fmt"
	"gin-api/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func IndexApi(c *gin.Context){
	var corpAdmin models.CorpAdminModel
	var account string
	account = "admin"
	models.Db().Where("account = ?", account).First(&corpAdmin)
	fmt.Println(corpAdmin.Salt)
	fmt.Println(corpAdmin.Password)
	c.String(http.StatusOK, "It works")
	c.String(http.StatusOK,"dealer !!")
}