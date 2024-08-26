package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Table struct {
	ID       int    `json:"id"`
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}

func GetAllTables(db *pgxpool.Pool) ([]Table, error) {
	rows, err := db.Query(context.Background(), "SELECT id, number, capacity, status FROM tables")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		err := rows.Scan(&table.ID, &table.Number, &table.Capacity, &table.Status)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}