package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-de-oliveira/crud-go/src/config/exception"
)

func CreateUser(c *gin.Context) {
	err := exception.NewBadRequestError("Method not Found")

	c.JSON(err.Code, err)
}
