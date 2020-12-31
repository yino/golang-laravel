package service

import (
	"go-laravel/app/models"
)

type RecordService struct {
}

func (p *RecordService) Page(params map[string]interface{}, page int, limit int) map[string]interface{} {
	// 总数
	var model models.MemberIntegralRecord
	db := models.Db().Limit(limit).Order("id desc")
	countDb := models.Db().Model(&model)
	memberId := params["memberId"].(int)
	if memberId > 0 {
		db = db.Where("member_id = ?", memberId)
		countDb = countDb.Where("member_id = ?", memberId)
	}
	var total int
	countDb.Count(&total)
	var list []models.MemberIntegralRecord
	// page limit
	page = page - 1
	if page > 0 {
		db = db.Offset(page * limit)
	}
	db.Find(&list)

	result := make(map[string]interface{})
	result["list"] = list
	result["meta"] = map[string]interface{}{
		"page":  page+1,
		"total": total,
		"size":  limit,
	}
	return result
}
