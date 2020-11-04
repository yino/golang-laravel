package models

import (
	"bytes"
	"fmt"
	"gin-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	_ "log"
	"time"
)

var (
	Db *gorm.DB
)

type Model struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}


// 数据库连接
func DbConnection() *gorm.DB {
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
	db, err := gorm.Open("mysql", connect.String())
	if err != nil {
		// 打日志
		log.Println("数据库连接错误----")
		fmt.Println("数据库连接错误----")
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
	}
	db.LogMode(true)
	defer db.Close()
	return db
}
