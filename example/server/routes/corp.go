package routes

import (
	"github.com/gin-gonic/gin"
	"go-laravel/app/http/controller/corp"
	"go-laravel/app/http/middleware"
)

func Corp(router *gin.Engine) {
	// 创建 corp 分组
	group := router.Group("/corp")
	group.POST("login", corp.LoginController)
	group.Use(middleware.CorpLoginAuth(), middleware.OperationRecords())
	{
		group.GET("get_info", corp.GetUserInfoController)
		group.GET("/index", corp.IndexController)

		// Member
		memberController := new(corp.MemberController)
		group.GET("/member/index", memberController.MemberIndex)
		group.POST("/member/add", memberController.MemberAdd)
		group.POST("/member/edit", memberController.MemberEdit)
		group.GET("/member/delete", memberController.MemberDelete)
		group.GET("/member/info", memberController.MemberInfo)
		group.POST("/member/add_integral", memberController.MemberIntegralAdd)

		// Product
		productController := new(corp.ProductController)
		group.GET("/product/all_list", productController.ProductList)
		group.GET("/product/index", productController.ProductIndex)
		group.POST("/product/add", productController.ProductAdd)
		group.POST("/product/edit", productController.ProductEdit)
		group.GET("/product/delete", productController.ProductDelete)
		group.GET("/product/info", productController.ProductInfo)
		group.POST("/product/upload", productController.ProductImgUpload)

		// Product
		recordController := new(corp.RecordController)
		group.GET("/record/index", recordController.RecordIndex)
	}
}
