package model

import (
	"fmt"

	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *exception.ExceptionErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
