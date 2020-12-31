package validate

type Product struct {
	Name string `form:"name" binding:"required"`
	Integral string `form:"integral" binding:"required"`
}
type ProductAdd struct {
	Name string `form:"name" binding:"required" json:"name"`
	Integral int `form:"integral" binding:"min=0" json:"integral"`
	Path string `form:"path" json:"path"`
}


type ProductEdit struct {
	Name string `form:"name" binding:"required" json:"name"`
	Integral int `form:"integral" binding:"min=0" json:"integral"`
	Path string `form:"path" json:"path"`
	Id int `form: "id" binding:"required"`
}