package models

import (
	"bytes"
	"fmt"
	"gin-api/config"
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// @title 公用模型
// @desc 常用字段
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

var DB *gorm.DB
// 数据库连接
func InitDbConnection() {

	fmt.Println("数据库连接池创建")
	mysqlConfig := config.MysqlConf()
	var connect bytes.Buffer
	connect.WriteString(mysqlConfig["User"].(string))
	connect.WriteString(":")
	connect.WriteString(mysqlConfig["Password"].(string))
	connect.WriteString("@/")
	connect.WriteString(mysqlConfig["Databases"].(string))
	connect.WriteString("?charset=")
	connect.WriteString(mysqlConfig["Charset"].(string))
	connect.WriteString("&parseTime=True&loc=Local")
	dbConnect,err := gorm.Open("mysql", connect.String())
	DB = dbConnect
	if err != nil {
		// 打日志
		log.Println("数据库连接错误----")
		fmt.Println("数据库连接错误----")
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
	}
	if config.AppDebug == true {
		DB.LogMode(true)
	}
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(mysqlConfig["MaxIdleConns"].(int))
	DB.DB().SetMaxOpenConns(mysqlConfig["MaxOpenConns"].(int))
	//defer DB.Close()
}

func Db() *gorm.DB {
	fmt.Println("--------get Db-------")
	//return InitDbConnection()
	fmt.Println(DB.DB().Ping())
	return DB
}
