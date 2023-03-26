package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"github.com/mateus-de-oliveira/crud-go/src/config/validation"
	"github.com/mateus-de-oliveira/crud-go/src/controller/models/request"
	"github.com/mateus-de-oliveira/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		exceptionErr := validation.ValidateUserError(err)

		c.JSON(exceptionErr.Code, exceptionErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Name, userRequest.Password, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
