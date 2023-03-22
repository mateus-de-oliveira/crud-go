package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-de-oliveira/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:userId", controller.FindUserById)
	r.GET("/users/email/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeleteUser)
}
