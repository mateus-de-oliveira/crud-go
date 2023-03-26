package model

import (
	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(string) *exception.ExceptionErr {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	return nil
}
