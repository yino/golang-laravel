package config

import (
	"reflect"
)

var (
	user = "root"
	password = "root"
	host = "127.0.0.1"
	databases = "go-gin-test"
	charset = "utf8"
	port = 3306
)

// 小写不能被访问转换
type mysqlConfig struct {
	User      string
	Password  string
	Host      string
	Databases string
	Charset   string
	Port      int
}

func MysqlConf() map[string]interface{} {
	config := mysqlConfig{
		User:      user,
		Password:  password,
		Host:      host,
		Databases: databases,
		Charset:   charset,
		Port:      port,
	}
	elem := reflect.ValueOf(&config).Elem()

	relType := elem.Type()
	data := make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}
