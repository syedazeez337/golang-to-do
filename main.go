package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/golang-to-do/config"
	"github.com/syedazeez337/golang-to-do/routes"
)

func main() {
	// Initialize dependencies
	config.InitDatabase()
	config.InitRedis()

	// Set up Gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
