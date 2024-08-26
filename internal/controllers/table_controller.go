package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kloudytics/restaurant-management-system/internal/models"
)

type TableController struct {
	DB *pgxpool.Pool
}

func NewTableController(db *pgxpool.Pool) *TableController {
	return &TableController{DB: db}
}

func (tc *TableController) GetAllTables(c *gin.Context) {
	tables, err := models.GetAllTables(tc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
		return
	}
	c.JSON(http.StatusOK, tables)
}

// Add more controller methods for CRUD operations on tables