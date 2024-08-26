package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kloudytics/restaurant-management-system/internal/models"
)

type OrderController struct {
	DB *pgxpool.Pool
}

func NewOrderController(db *pgxpool.Pool) *OrderController {
	return &OrderController{DB: db}
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := models.GetAllOrders(oc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Add more controller methods for CRUD operations on orders