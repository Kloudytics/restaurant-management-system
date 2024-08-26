package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Order struct {
	ID          int       `json:"id"`
	TableID     int       `json:"table_id"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	ID         int     `json:"id"`
	OrderID    int     `json:"order_id"`
	MenuItemID int     `json:"menu_item_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

func GetAllOrders(db *pgxpool.Pool) ([]Order, error) {
	rows, err := db.Query(context.Background(), `
		SELECT o.id, o.table_id, o.status, o.total_amount, o.created_at,
			   oi.id, oi.menu_item_id, oi.quantity, oi.price
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		ORDER BY o.id, oi.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make(map[int]*Order)
	for rows.Next() {
		var o Order
		var oi OrderItem
		err := rows.Scan(&o.ID, &o.TableID, &o.Status, &o.TotalAmount, &o.CreatedAt,
			&oi.ID, &oi.MenuItemID, &oi.Quantity, &oi.Price)
		if err != nil {
			return nil, err
		}

		if order, exists := orders[o.ID]; exists {
			order.Items = append(order.Items, oi)
		} else {
			o.Items = []OrderItem{oi}
			orders[o.ID] = &o
		}
	}

	result := make([]Order, 0, len(orders))
	for _, order := range orders {
		result = append(result, *order)
	}

	return result, nil
}

// Add more functions for CRUD operations on orders