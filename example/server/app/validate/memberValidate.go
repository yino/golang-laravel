package validate

type Member struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Mobile string `form:"mobile" binding:"required" json:"mobile"`
}
type MemberAdd struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Mobile   string `form:"mobile" binding:"required" json:"mobile"`
	Integral int    `form:"integral" binding:"min=0" json:"integral"`
}

type MemberEdit struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Mobile   string `form:"mobile" binding:"required" json:"mobile"`
	Integral int    `form:"integral" binding:"min=0" json:"integral"`
	Id       int    `form: "id" binding:"required" json:"id"`
}

type MemberAddIntegral struct {
	MemberId  int `form:"member_id" binding:"required" json:"member_id"`
	ProductId int `form:"product_id" binding:"required" json:"product_id"`
	Integral  int `form:"integral" binding:"required" json:"integral"`
	Money  int `form:"money" binding:"min=0" json:"money"`
}
