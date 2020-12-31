package models

import (
	"bytes"
	"fmt"
	"go-laravel/config"
	"go-laravel/extend/log"
	"strconv"
	"time"

	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// @title 公用模型
// @desc 常用字段
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt int        `json:"created_at"`
	UpdatedAt int        `json:"updated_at"`
	DeletedAt *time.Time `json:"-",sql:"index"`
}

var DB *gorm.DB

// InitDbConnection db connection pool
func InitDbConnection(pool bool) {
	log.Println("------mysql database connection")
	mysqlConfig := config.MysqlConf()
	var connect bytes.Buffer
	connect.WriteString(mysqlConfig["User"].(string))
	connect.WriteString(":")
	connect.WriteString(mysqlConfig["Password"].(string))
	connect.WriteString("@tcp(")
	connect.WriteString(mysqlConfig["Host"].(string))
	connect.WriteString(":")
	connect.WriteString(strconv.Itoa(mysqlConfig["Port"].(int)))
	connect.WriteString(")")
	connect.WriteString("/")
	connect.WriteString(mysqlConfig["Databases"].(string))
	connect.WriteString("?charset=")
	connect.WriteString(mysqlConfig["Charset"].(string))
	connect.WriteString("&parseTime=True&loc=Local")
	fmt.Println(connect.String())
	dbConnect, err := gorm.Open("mysql", connect.String())
	DB = dbConnect

	//
	DB.Callback().Update().Register("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
		}
	})
	DB.Callback().Create().Register("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if !scope.HasError() {
			if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
				if createdAtField.IsBlank {
					_ = createdAtField.Set(time.Now().Unix())
				}
			}
			if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
				if updatedAtField.IsBlank {
					_ = updatedAtField.Set(time.Now().Unix())
				}
			}
		}
	})

	if err != nil {
		log.SetPrefix("error")
		log.Println("mysql connect fail----", err)
	} else {
		log.SetPrefix("【success】")
		log.Println("Mysql connection success")
	}
	if config.AppDebug || config.DatabaseDebug {
		DB.LogMode(true)
	}

	if pool {
		DB.SingularTable(true)
		DB.DB().SetMaxIdleConns(mysqlConfig["MaxIdleCons"].(int))
		DB.DB().SetMaxOpenConns(mysqlConfig["MaxOpenCons"].(int))
		DB.DB().SetConnMaxLifetime(time.Second * 60)
	}
	//defer DB.Close()
}

// Db conncet
func Db() *gorm.DB {
	if err := DB.DB().Ping(); err != nil {
		//DB.Close()
		log.Println(err)
		log.Fatal("-------------数据库ping error-------------------")
		InitDbConnection(true)
	}
	return DB
}
