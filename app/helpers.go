package app

import (
	"fmt"
	"go-laravel/config"
	"os"
)

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
// 判断文件夹是否存在
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil{
		return true
	}
	return false
}

func GetLogDir() string{
	dir, _ := os.Getwd()
	return dir+ config.LogPath
}
