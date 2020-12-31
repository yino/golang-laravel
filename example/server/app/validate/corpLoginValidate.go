package validate

type CorpLogin struct {
	Account string `form:"account" binding:"required" json:"account"`
	Password string `form:"password" binding:"required" json:"password"`
}