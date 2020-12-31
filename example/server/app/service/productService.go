package service

import (
	"errors"
	"go-laravel/app/models"
	"go-laravel/config"
)

type ProductService struct {
}

func (p *ProductService) List() []models.ProductModel {
	var list []models.ProductModel
	models.Db().Order("id desc").Find(&list)
	return list
}
func (p *ProductService) Page(params map[string]interface{}, page int, limit int) map[string]interface{} {
	// 总数
	db := models.Db().Limit(limit).Order("id desc")
	countDb := models.Db().Model(&models.ProductModel{})
	if len(params["name"].(string)) > 0 {
		name := params["name"].(string)
		name = "%"+name+"%"
		db = db.Where("name like ?", name)
		countDb = countDb.Where("name like ?", name)
	}
	var total int
	countDb.Count(&total)
	var list []models.ProductModel
	// page limit
	page = page - 1
	if page > 0 {
		db = db.Offset(page * limit)
	}
	db.Find(&list)

	result := make(map[string]interface{})
	result["list"] = list
	result["meta"] = map[string]interface{}{
		"page":  page+1,
		"total": total,
		"size":  limit,
	}
	return result
}

func (p *ProductService) Add(params map[string]interface{}) error {

	// 检查名称是否存在
	var data models.ProductModel
	models.Db().Where("name = ?", params["name"].(string)).First(&data)

	if data.ID > 0 {
		return errors.New("名称已存在")
	}
	model := models.ProductModel{
		Name:     params["name"].(string),
		Integral: params["integral"].(int),
		Image: params["path"].(string),
	}
	result := models.Db().Create(&model)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (p *ProductService) Info(id int) (interface{}, error) {
	var data models.ProductModel
	models.Db().Where("id = ?", id).First(&data)
	if data.ID == 0 {
		return false, errors.New("id 不存在")
	}

	result := make(map[string]interface{})
	result["id"] = data.ID
	result["name"] = data.Name
	result["path"] = data.Image
	if len(data.Image) >0 {
		result["remote_url"] = "http://"+config.AppPort+data.Image
	}else{
		result["remote_url"] = ""
	}

	result["integral"] = data.Integral
	return result, nil
}

func (p *ProductService) Edit(params map[string]interface{}) error {
	var data models.ProductModel
	// 检查name是否存在
	models.Db().
		Where("name = ?", params["name"].(string)).
		Where("id <> ?", params["id"].(int)).
		First(&data)
	if data.ID > 0 {
		return errors.New("name已存在")
	}
	// 查询当前数据
	models.Db().
		Model(models.ProductModel{}).
		Where("id = ?", params["id"].(int)).
		First(&data)
	data.Name = params["name"].(string)
	data.Integral = params["integral"].(int)
	data.Image = params["image"].(string)

	res := models.Db().Save(&data)
	return res.Error
}

func (p *ProductService) Delete(id int) error {
	var data models.ProductModel
	err := models.Db().Where("id = ?", id).First(&data).Error
	if err != nil {
		return errors.New("此数据 不存在")
	}
	res := models.Db().Where("id=?", id).Delete(&data)
	return res.Error
}
