package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
)

// @title t_corp_admin
type MemberModel struct {
	Model
	Name     string `gorm:"default:''" json:"name"`
	Mobile   string `gorm:"default:''" json:"mobile"`
	Integral int    `gorm:"default:0" json:"integral"`
	CompanyId int `gor`
}

func (MemberModel) TableName() string {
	return "sys_member"
}

const ()
