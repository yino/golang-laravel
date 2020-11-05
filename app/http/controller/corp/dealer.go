package corp

import (
	"gin-api/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(c *gin.Context) {
	var corpAdminModel []models.CorpAdminModel
	models.Db().Find(corpAdminModel)

	c.String(http.StatusOK, "It works")
}



