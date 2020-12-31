package models

// @title t_corp_admin
type ProductModel struct {
	Model
	Name     string `gorm:"default:''" json:"name"`
	Image   string `gorm:"default:''" json:"image"`
	Integral int `gorm:"default:0" json:"integral"`
}

func (ProductModel) TableName() string {
	return "sys_product"
}
