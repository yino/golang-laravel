package corp

import (
	"github.com/gin-gonic/gin"
	response "go-laravel/app/exception"
	"go-laravel/app/service"
)

func IndexController(ctx * gin.Context)  {

	cbs := service.CorpCount{}
	// 统计 多少用户
	var countService service.Count
	countService = cbs
	memberCount := countService.MemberCount()
	// 统计 多少推本
	productCount := countService.ProductCount()
	//
	moneySum := countService.MoneySum()
	recordCount := countService.RecordCount()
	weekGroupCount := countService.WeekGroupCount()
	data := make(map[string]interface{})
	data["member_count"] = memberCount
	data["product_count"] = productCount
	data["total_money"] = moneySum
	data["count"] = recordCount
	data["week_group_count"] = weekGroupCount
	response.SuccessResponse(data)
}