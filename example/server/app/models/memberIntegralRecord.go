package models

// member integral record
type MemberIntegralRecord struct {
	Model
	MemberMobile    string `gorm:"default:''" json:"member_mobile"`
	MemberName      string `gorm:"default:''" json:"member_name"`
	Integral        int    `gorm:"default:0'" json:"integral"`
	SurplusIntegral int    `gorm:"default:0'" json:"surplus_integral"`
	MemberId        int    `gorm:"default:0'" json:"member_id"`
	ProductId       int    `gorm:"default:0'" json:"product_id"`
	ProductName     string `gorm:"default:''" json:"product_name""`
	UserId          int    `gorm:"default:0" json:"user_id"`
	UserName        string `gorm:"default:''" json:"user_name"`
	Money           int    `gorm:"default:0" json:"money"`
	Date            string `gorm:"default:''" json:"date"`
}

type MoneySum struct {
	TotalMoney int `json:"total_money"`
}

type GroupByCount struct {
	GroupCount int    `json:"group_count"`
	Date       string `json:"date"`
}

func (MemberIntegralRecord) TableName() string {
	return "sys_member_integral_record"
}
