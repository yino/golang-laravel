package validate

type CorpLogin struct {
	Account string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}