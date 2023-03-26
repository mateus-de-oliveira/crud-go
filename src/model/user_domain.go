package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	FindUser(string) *exception.ExceptionErr
	CreateUser() *exception.ExceptionErr
	UpdateUser(string) *exception.ExceptionErr
	DeleteUser(string) *exception.ExceptionErr
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()
	hash.Write([]byte(ud.Password))

	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
