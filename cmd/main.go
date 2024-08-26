package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kloudytics/restaurant-management-system/internal/database"
	"github.com/kloudytics/restaurant-management-system/internal/routes"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}