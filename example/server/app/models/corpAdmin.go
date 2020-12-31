package models

// CorpAdminModel table=sys_corp_admin
type CorpAdminModel struct {
	Model
	Name          string `gorm:"default:''"`
	Account       string `gorm:"default:''"`
	Password      string `gorm:"default:''"`
	Salt          string
	LastLoginTime int
	Status        int `gorm:"default:1"`
	Token         string
	CompanyID     int `gorm:"deafult:0" json:"company_id"`
}

// TableName set
func (CorpAdminModel) TableName() string {
	return "sys_corp_admin"
}

const (
	// 禁用状态
	CorpAdminStatusClose = 2
)
