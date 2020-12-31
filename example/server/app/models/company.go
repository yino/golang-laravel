package models

// @title sys_company
type CompanyModel struct {
	Model
	CompanyName string `gorm:"default:''" json:"company_name"`
}

func (CompanyModel) TableName() string {
	return "sys_company"
}
