package config

import (
	"reflect"
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
		User:      "root",
		Password:  "root",
		Host:      "127.0.0.1",
		Databases: "b2b_dev_common",
		Charset:   "utf8",
		Port:      3306,
	}
	elem := reflect.ValueOf(&config).Elem()

	relType := elem.Type()
	data := make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}
