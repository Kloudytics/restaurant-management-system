package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kloudytics/restaurant-management-system/internal/controllers"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	menuController := controllers.NewMenuController(db)
	tableController := controllers.NewTableController(db)
	orderController := controllers.NewOrderController(db)

	v1 := router.Group("/api/v1")
	{
		// Menu routes
		v1.GET("/menu", menuController.GetAllMenuItems)
		v1.POST("/menu", menuController.CreateMenuItem)
		v1.GET("/menu/:id", menuController.GetMenuItem)
		v1.PUT("/menu/:id", menuController.UpdateMenuItem)
		v1.DELETE("/menu/:id", menuController.DeleteMenuItem)

		// Table routes
		v1.GET("/tables", tableController.GetAllTables)

		// Order routes
		v1.GET("/orders", orderController.GetAllOrders)
	}
}