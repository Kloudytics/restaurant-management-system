package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func GetAllMenuItems(db *pgxpool.Pool) ([]MenuItem, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name, description, price, category FROM menu_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MenuItem
	for rows.Next() {
		var item MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Category)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Add more functions for CRUD operations as needed