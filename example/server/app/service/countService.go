package service

import (
	"fmt"
	"go-laravel/app/models"
	"strings"
	"time"
)

type Count interface {
	MemberCount() int
	ProductCount() int
	MoneySum() int
	RecordCount() int
	WeekGroupCount() []map[string]interface{}
}
type CorpCount struct {
}

func (c CorpCount) MemberCount() int {
	var count int
	models.Db().Model(&models.MemberModel{}).Count(&count)
	return count
}

func (c CorpCount) ProductCount() int {
	var count int
	models.Db().Model(&models.ProductModel{}).Count(&count)
	return count
}

func (c CorpCount) MoneySum() int {

	//var count []int
	var result models.MoneySum
	models.Db().
		Table("sys_member_integral_record").
		Select("SUM(`money`) as total_money").
		Scan(&result)
	return result.TotalMoney
}

func (c CorpCount) RecordCount() int {
	var count int
	models.Db().Model(&models.MemberIntegralRecord{}).Count(&count)
	return count
}

func (c CorpCount) WeekGroupCount() []map[string]interface{} {

	var date string
	dateArr := make([]string, 7)
	//dateArr = "("
	for i := 0; i < 7; i++ {
		if i == 0 {
			date = time.Now().Format("2006-01-02")
		} else {
			date = time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		}
		dateArr[i] = date
	}
	var result []models.GroupByCount
	var model models.MemberIntegralRecord
	models.Db().
		Table(model.TableName()).
		Select("count(*) as group_count, `date`").
		Where("`date` IN (?)", dateArr).
		Order("`date` Desc").
		Group("`date`").
		Scan(&result)

	arrResult := make([]map[string]interface{}, 7)
	mapResult := make(map[string]int)
	for _, val := range result {
		test := strings.Split(val.Date,"T")
		date = test[0]
		mapResult[date] = val.GroupCount
	}

	//index := 0

	for k, v:= range dateArr{
		item := make(map[string]interface{})
		item["name"] = v
		if _, ok := mapResult[v]; !ok {
			item["value"] = 0
			//mapResult[v] = 0
		} else{
			item["value"] = mapResult[v]
		}

		fmt.Println(k, item)
		arrResult[k] = item
		//index = index+1
	}
	fmt.Println(arrResult)
	return arrResult
}
