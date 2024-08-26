package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kloudytics/restaurant-management-system/internal/models"
)

type MenuController struct {
	DB *pgxpool.Pool
}

func NewMenuController(db *pgxpool.Pool) *MenuController {
	return &MenuController{DB: db}
}

func (mc *MenuController) GetAllMenuItems(c *gin.Context) {
	items, err := models.GetAllMenuItems(mc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Add more controller methods as needed