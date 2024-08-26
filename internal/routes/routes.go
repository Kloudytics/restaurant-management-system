package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kloudytics/restaurant-management-system/internal/controllers"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	menuController := controllers.NewMenuController(db)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/menu", menuController.GetAllMenuItems)
		// Add more routes as needed
	}
}