package model

import (
	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) DeleteUser(string) *exception.ExceptionErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
