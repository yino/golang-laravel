package service

import (
	"go-laravel/app/models"
)

type OperationRecordsService struct {
}

// 添加日志
func (m *OperationRecordsService) Add(params map[string]interface{}) {
	model := models.OperationRecordsModel{
		Ip:      params["ip"].(string),
		Method:  params["method"].(string),
		Resp:    params["resp"].(string),
		Path:    params["path"].(string),
		Status:  params["status"].(int),
		Latency: params["latency"].(int),
		Agent:   params["agent"].(string),
		//ErrorMessage: params["error_message"].(string),
		Body:      params["body"].(string),
		UserId:    params["userId"].(int),
		StartTime: params["startTime"].(int),
		EndTime:   params["endTime"].(int),
	}
	models.Db().Create(&model)
}

// 查看日志
func (m *OperationRecordsService) Page() {

}
