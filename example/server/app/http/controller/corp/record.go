package corp

import (
	"github.com/gin-gonic/gin"
	response "go-laravel/app/exception"
	"go-laravel/app/service"
	"go-laravel/extend/log"
	"strconv"
)

type Record interface {
	RecordIndex()
}
type RecordController struct {
	ServiceObj *service.RecordService
}
func (m *RecordController) RecordIndex(request *gin.Context) {
	page := request.DefaultQuery("page", "1")
	size := request.DefaultQuery("size", "10")
	memberId := request.DefaultQuery("member_id", "0")
	log.Write("page：", page)
	// 参数拼接
	params := make(map[string]interface{})


	// 分页
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	memberIdInt, _ := strconv.Atoi(memberId)
	params["memberId"] = memberIdInt

	result := m.ServiceObj.Page(params, pageInt, sizeInt)
	response.SuccessResponse(result)
}