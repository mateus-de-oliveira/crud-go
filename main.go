package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-de-oliveira/crud-go/src/config/logger"
	"github.com/mateus-de-oliveira/crud-go/src/controller/routes"
)

const (
	PORT_ENV = "PORT"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading.env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	logger.Info("Application is running on port: " + os.Getenv(PORT_ENV))

	if err := router.Run(":" + os.Getenv(PORT_ENV)); err != nil {
		log.Fatal(err)
	}

}
