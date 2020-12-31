package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
)

type OperationRecordsModel struct {
	Model
	Ip           string `gorm:"default:''" json:"ip"`
	Method       string `gorm:"default:''" json:"method"`
	Path         string `gorm:"default:0" json:"path"`
	Status       int    `gorm:"default:0" json:"status"`
	Latency      int    `gorm:"default:0" json:"latency"`
	Agent        string `gorm:"default:0" json:"agent"`
	ErrorMessage string `gorm:"default:0" json:"error_message"`
	Body         string `gorm:"default:0" json:"body"`
	Resp         string `gorm:"default:0" json:"resp"`
	UserId       int    `gorm:"default:0" json:"user_id"`
	StartTime    int    `gorm:"default:0" json:"start_time"`
	EndTime      int    `gorm:"default:0" json:"end_time"`
}

func (OperationRecordsModel) TableName() string {
	return "sys_operation_records"
}
