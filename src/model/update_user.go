package model

import (
	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(string) *exception.ExceptionErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
