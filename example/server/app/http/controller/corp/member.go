package corp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	response "go-laravel/app/exception"
	"go-laravel/app/service"
	"go-laravel/app/validate"
	"strconv"
)

type Member interface {
	MemberList()
	MemberAdd()
	MemberEdit()
	MemberDelete()
	MemberAddIntegral()
}

type MemberController struct {
	ServiceObj *service.MemberService
}

func (m *MemberController) MemberList(request *gin.Context) {
	m.ServiceObj.List()
	response.SuccessResponse("")
}

func (m *MemberController) MemberIndex(request *gin.Context) {
	page := request.DefaultQuery("page", "1")
	size := request.DefaultQuery("size", "10")
	mobile := request.Query("mobile")
	name := request.Query("name")
	// 参数拼接
	params := make(map[string]interface{})
	params["mobile"] = mobile
	params["name"] = name

	// 分页
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	result := m.ServiceObj.Page(params, pageInt, sizeInt)

	response.SuccessResponse(result)
}

func (m *MemberController) MemberAdd(request *gin.Context) {
	var validates validate.MemberAdd
	if err := request.ShouldBindJSON(&validates); err != nil {
		fmt.Println(validates)
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		validate.ErrResp(errs)
		//response.BaseException(500000, errs.Error())
	}
	name := validates.Name
	mobile := validates.Mobile
	integral := validates.Integral
	params := make(map[string]interface{})
	params["name"] = name
	params["mobile"] = mobile
	params["integral"] = integral

	res := m.ServiceObj.Add(params)

	if res != nil {
		response.BaseException(5000, res.Error())
	} else {
		response.SuccessResponse("添加成功")
	}
}

func (m *MemberController) MemberEdit(request *gin.Context) {
	var validates validate.MemberEdit
	if err := request.ShouldBindJSON(&validates); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		validate.ErrResp(errs)
		//response.BaseException(500000, errs.Error())
	}

	formId := validates.Id
	formName := validates.Name
	formMobile := validates.Mobile
	formIntegral := validates.Integral
	fmt.Println(formId)
	params := make(map[string]interface{})

	params["integral"] = formIntegral
	params["id"] = formId
	params["name"] = formName
	params["mobile"] = formMobile

	err := m.ServiceObj.Edit(params)
	if err != nil {
		response.BaseException(50000, err.Error())
	} else {
		response.SuccessResponse("修改成功")
	}
}

func (m *MemberController) MemberInfo(request *gin.Context) {
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

func (m *MemberController) MemberDelete(request *gin.Context) {
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

// 增加用户积分
func (m *MemberController) MemberIntegralAdd(request *gin.Context) {

	var validates validate.MemberAddIntegral
	if err := request.ShouldBindJSON(&validates); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		// validator.ValidationErrors类型错误则进行翻译
		validate.ErrResp(errs)
	}

	memberId := validates.MemberId
	integral := validates.Integral
	productId := validates.ProductId
	money := validates.Money*100

	params := make(map[string]interface{})
	params["memberId"] = memberId
	params["integral"] = integral
	params["productId"] = productId
	params["money"] = money
	corpInfo, _ := request.Get("CorpInfo")
	CorpInfo := corpInfo.(map[string]interface{})
	params["userId"] = CorpInfo["id"].(int)
	params["userName"] = CorpInfo["username"].(string)
	params["userName"] = CorpInfo["username"].(string)
	err := m.ServiceObj.AddIntegral(params)
	if err != nil {
		response.BaseException(500000, err.Error())
	} else {
		response.SuccessResponse("")
	}
}
