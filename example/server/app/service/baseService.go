package service

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"go-laravel/config"
)


type Operation interface {
	List()
	Page()
	Add()
	Info()
	Edit()
	Delete()
}

type BaseService struct {

}

// param id int 用户id
// param userType string 用户类型 corp customer admin
func (b *BaseService) CreateUserJwtToken(uid string, userType string)  (string, error){
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"exp":  time.Now().Add(config.LoginTokenExpire).Unix(),
	})
	token, err := at.SignedString([]byte(userType))
	if err != nil {
		return "", err
	}
	return token, nil
}