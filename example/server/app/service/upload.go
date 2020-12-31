package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-laravel/config"
	"mime/multipart"
	"os"
	"reflect"
	"time"
)

type Upload struct {
}

func (u *Upload) Img(file *multipart.FileHeader, dir string, c *gin.Context) (string, error) {
	//获取文件后缀
	extString := Ext(file.Filename)
	var err error
	if extString == "" {
		err = errors.New("上传失败，文件类型不支持，只能上传 图片。")
		return "", err
	}

	extStrSlice := []string{".png", ".jpg", ".JPG", ".PNG", ".JPEG"}
	if !ContainArray(extString, extStrSlice) {
		//fmt.Println("上传失败，文件类型不支持，只能上传 xlsx 格式的。")
		err = errors.New("上传失败，文件类型不支持，只能上传.png, jpg, jpg, .PNG, .JPEG 格式的。")
		return "", err
	}

	//filepath := 'resource/public/uploads/'
	filepath := config.FILE_ROOT_PATH + dir //从配置文件里取
	//如果没有filepath文件目录就创建一个
	_, err = os.Stat(filepath)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(filepath, os.ModePerm)
			fmt.Println("创建文件 err：", err)
		}
	}
	//上传到的路径
	//path := 'resource/public/uploads/20060102150405test.xlsx'
	file.Filename = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	filepath = filepath+"/" + file.Filename                                                        //路径+文件名上传

	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		err = errors.New(fmt.Sprintf("上传失败，%v", err))
		return "", err
	}
	filepath = filepath[1:len(filepath)]

	return filepath , nil

}
func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

//Contain 判断obj是否在target中，target支持的类型array,slice,map   false:不在 true:在
func ContainArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}
