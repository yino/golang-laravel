package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-laravel/app/models"
	"time"
)

type MemberService struct {
}

func (m *MemberService) List() {

}

func (m *MemberService) Page(params map[string]interface{}, page int, limit int) map[string]interface{} {
	var model models.MemberModel
	// 总数
	db := models.Db().Limit(limit).Order("id desc")
	countDb := models.Db().Model(&model)
	if len(params["mobile"].(string)) > 0 {
		mobile := params["mobile"].(string)
		db = db.Where("mobile = ?", mobile)
		countDb = countDb.Where("mobile = ?", mobile)
	}
	if len(params["name"].(string)) > 0 {
		name := params["name"].(string)
		name = "%"+name+"%"
		db = db.Where("`name` like ?", name)
		countDb = countDb.Where("`name` like ? ", name)
	}
	var total int
	countDb.Count(&total)
	var list []models.MemberModel
	// page limit
	page = page - 1
	if page > 0 {
		db = db.Offset(page * limit)
	}

	db.Find(&list)

	result := make(map[string]interface{})
	result["list"] = list
	result["meta"] = map[string]interface{}{
		"page":  page + 1,
		"total": total,
		"size":  limit,
	}
	return result
}

func (m *MemberService) Add(params map[string]interface{}) error {

	// 检查手机号是否存在
	var data models.MemberModel
	models.Db().Where("mobile = ?", params["mobile"].(string)).First(&data)

	if data.ID > 0 {
		return errors.New("手机号已存在")
	}
	model := models.MemberModel{
		Name:     params["name"].(string),
		Integral: params["integral"].(int),
		Mobile:   params["mobile"].(string),
	}
	result := models.Db().Create(&model)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (m *MemberService) Info(id int) (interface{}, error) {
	var data models.MemberModel
	models.Db().Where("id = ?", id).First(&data)
	if data.ID == 0 {
		return false, errors.New("id 不存在")
	}
	return data, nil
}

func (m *MemberService) Edit(params map[string]interface{}) error {
	var data models.MemberModel
	// 检查电话号码是否存在
	models.Db().
		Where("mobile = ?", params["mobile"].(string)).
		Where("id <> ?", params["id"].(int)).
		First(&data)
	if data.ID > 0 {
		return errors.New("数据已存在")
	}
	// 查询当前数据
	models.Db().
		Model(models.MemberModel{}).
		Where("id = ?", params["id"].(int)).
		First(&data)
	data.Name = params["name"].(string)
	data.Mobile = params["mobile"].(string)
	data.Integral = params["integral"].(int)

	res := models.Db().Save(&data)
	return res.Error
}

func (m *MemberService) Delete(id int) error {
	var data models.MemberModel
	err := models.Db().Where("id = ?", id).First(&data).Error
	if err != nil {
		return errors.New("此数据 不存在")
	}
	res := models.Db().Where("id=?", id).Delete(&data)
	return res.Error
}

// @param params map {"memberId" int , "productId" int, "integral" int}
func (m *MemberService) AddIntegral(params map[string]interface{}) error {

	// 查询member id 是否存在
	var memberInfo models.MemberModel
	var productInfo models.ProductModel
	models.Db().Where("id = ?", params["memberId"]).First(&memberInfo)
	if memberInfo.ID == 0 {
		return errors.New("用户不存在")
	}
	// 查询 product id 是否存在
	models.Db().Where("id = ?", params["productId"]).First(&productInfo)
	if productInfo.ID == 0 {
		return errors.New("此推本不存在")
	}

	member := models.MemberModel{}
	product := models.MemberIntegralRecord{
		MemberMobile:    memberInfo.Mobile,
		MemberName:      memberInfo.Name,
		Integral:        params["integral"].(int),
		SurplusIntegral: memberInfo.Integral,
		MemberId:        params["memberId"].(int),
		ProductId:       params["productId"].(int),
		ProductName:     productInfo.Name,
		UserId:          params["userId"].(int),
		UserName:        params["userName"].(string),
		Money:           params["money"].(int),
		Date:            time.Unix(time.Now().Unix(), 0).Format("2006-01-02"),
	}

	err := models.Db().Transaction(func(tx *gorm.DB) error {
		var err error
		// 积分修改
		err = tx.Model(&member).
			Where("id = ?", params["memberId"]).
			Update("integral", gorm.Expr("integral +?", params["integral"])).Error
		if err != nil {
			return err
		}
		// 积分日志增加
		err = tx.Create(&product).Error
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
