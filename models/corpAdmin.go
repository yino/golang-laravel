package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
	"time"
)

type CorpAdminModel struct {
	Model
	Username      string `gorm:"default:''"`
	Account       string `gorm:"default:''"`
	Password      string `gorm:"default:''"`
	Salt          string
	LastLoginTime *time.Location
	Status        int `gorm:"default:1"`
}

func (CorpAdminModel) TableName() string {
	return "t_corp_admin"
}

