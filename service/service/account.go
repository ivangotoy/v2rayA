package service

import (
	"V2RayA/persistence/configure"
	"V2RayA/common"
	"V2RayA/common/jwt"
	"errors"
	"time"
)

func Login(username, password string) (token string, err error) {
	if !IsValidAccount(username, password) {
		return "", errors.New("invalid username or password")
	}
	dur := 3 * time.Hour
	return jwt.MakeJWT(map[string]string{
		"uname": username,
	}, &dur)
}

func IsValidAccount(username, password string) bool {
	pwd, err := configure.GetPasswordOfAccount(username)
	if err != nil {
		return false
	}
	return pwd == common.CryptoPwd(password)
}

func Register(username, password string) (token string, err error) {
	if configure.ExistsAccount(username) {
		return "", errors.New("username exists")
	}
	err = configure.SetAccount(username, common.CryptoPwd(password))
	if err != nil {
		return
	}
	return Login(username,password)
}

func ValidPasswordLength(password string) bool {
	return len(password) >= 5 && len(password) <= 32
}
