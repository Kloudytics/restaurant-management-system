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

func CreateMenuItem(db *pgxpool.Pool, item *MenuItem) error {
	_, err := db.Exec(context.Background(),
		"INSERT INTO menu_items (name, description, price, category) VALUES ($1, $2, $3, $4)",
		item.Name, item.Description, item.Price, item.Category)
	return err
}

func GetMenuItem(db *pgxpool.Pool, id int) (MenuItem, error) {
	var item MenuItem
	err := db.QueryRow(context.Background(),
		"SELECT id, name, description, price, category FROM menu_items WHERE id = $1", id).
		Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Category)
	return item, err
}

func UpdateMenuItem(db *pgxpool.Pool, item *MenuItem) error {
	_, err := db.Exec(context.Background(),
		"UPDATE menu_items SET name = $1, description = $2, price = $3, category = $4 WHERE id = $5",
		item.Name, item.Description, item.Price, item.Category, item.ID)
	return err
}

func DeleteMenuItem(db *pgxpool.Pool, id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM menu_items WHERE id = $1", id)
	return err
}