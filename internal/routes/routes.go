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
		v1.POST("/menu", menuController.CreateMenuItem)
		v1.GET("/menu/:id", menuController.GetMenuItem)
		v1.PUT("/menu/:id", menuController.UpdateMenuItem)
		v1.DELETE("/menu/:id", menuController.DeleteMenuItem)
	}
}