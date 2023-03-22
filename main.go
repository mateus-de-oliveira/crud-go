package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-de-oliveira/crud-go/src/controller/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading.env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}

}
