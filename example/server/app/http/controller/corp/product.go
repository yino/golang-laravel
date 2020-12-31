package corp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	response "go-laravel/app/exception"
	"go-laravel/app/service"
	"go-laravel/app/validate"
	"go-laravel/config"
	"go-laravel/extend/log"
	"strconv"
)

type Product interface {
	ProductList()
	ProductAdd()
	ProductEdit()
	ProductDelete()
	ProductImgUpload()
}

type ProductController struct {
	ServiceObj *service.ProductService
}

func (m *ProductController) ProductList(request *gin.Context) {
	result := m.ServiceObj.List()
	response.SuccessResponse(result)
}

func (m *ProductController) ProductIndex(request *gin.Context) {
	page := request.DefaultQuery("page", "1")
	size := request.DefaultQuery("size", "10")
	name := request.Query("name")
	log.Write("page：", page)
	// 参数拼接
	params := make(map[string]interface{})
	params["name"] = name

	// 分页
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	result := m.ServiceObj.Page(params, pageInt, sizeInt)

	response.SuccessResponse(result)
}

func (m *ProductController) ProductAdd(request *gin.Context) {

	var validates validate.ProductAdd
	if err := request.ShouldBindJSON(&validates); err != nil {
		fmt.Println(err)
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		validate.ErrResp(errs)
	}
	name := validates.Name
	integral := validates.Integral
	path := validates.Path
	params := make(map[string]interface{})

	params["name"] = name
	params["integral"] = integral
	params["path"] = path
	res := m.ServiceObj.Add(params)
	if res != nil {
		response.BaseException(5000, res.Error())
	} else {
		response.SuccessResponse("添加成功")
	}
}

func (m *ProductController) ProductEdit(request *gin.Context) {
	var validates validate.ProductEdit
	if err := request.ShouldBindJSON(&validates); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		validate.ErrResp(errs)
		//response.BaseException(500000, errs.Error())
	}

	id := validates.Id
	name := validates.Name
	path := validates.Path
	integral := validates.Integral

	params := make(map[string]interface{})
	var err error
	params["id"] = id
	params["name"] = name
	params["image"] = path
	params["integral"] = integral

	err = m.ServiceObj.Edit(params)
	if err != nil {
		response.BaseException(50000, err.Error())
	} else {
		response.SuccessResponse("修改成功")
	}
}

func (m *ProductController) ProductInfo(request *gin.Context) {
	queryId, queryRes := request.GetQuery("id")
	if !queryRes {
		response.ValidateResponse("请输入id")
	}
	id, err := strconv.Atoi(queryId)
	if err != nil || id <= 0 {
		response.ValidateResponse("id 参数错误")
	}

	var info interface{}
	info, err = m.ServiceObj.Info(id)
	if err != nil {
		response.BaseException(500000, err.Error())
	}
	response.SuccessResponse(info)
}

func (m *ProductController) ProductDelete(request *gin.Context) {
	queryId, queryRes := request.GetQuery("id")
	if !queryRes {
		response.ValidateResponse("请输入id")
	}
	id, err := strconv.Atoi(queryId)
	if err != nil || id <= 0 {
		response.ValidateResponse("id 参数错误")
	}
	err = m.ServiceObj.Delete(id)
	if err != nil {
		response.BaseException(500000, err.Error())
	} else {
		response.SuccessResponse("删除成功")
	}
}
func (m *ProductController) ProductImgUpload(request *gin.Context) {
	f, err := request.FormFile("img")
	//错误处理
	if err != nil {
		response.BaseException(50000, err.Error())
	} else {
		upload := new(service.Upload)
		path, err := upload.Img(f, "product", request)
		if err != nil {
			response.BaseException(500000, err.Error())
		} else {
			result := make(map[string]string)
			fmt.Println("file domain：", config.FILE_DOMAON+path)
			result["remote_url"] = config.FILE_DOMAON + path
			result["path"] = path
			response.SuccessResponse(result)
		}

	}
}
