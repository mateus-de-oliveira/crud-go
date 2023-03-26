package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"github.com/mateus-de-oliveira/crud-go/src/config/validation"
	"github.com/mateus-de-oliveira/crud-go/src/controller/models/request"
	"github.com/mateus-de-oliveira/crud-go/src/controller/models/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		exceptionErr := validation.ValidateUserError(err)

		c.JSON(exceptionErr.Code, exceptionErr)
		return
	}

	response := response.UserResponse{
		ID:    "TESTE",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
