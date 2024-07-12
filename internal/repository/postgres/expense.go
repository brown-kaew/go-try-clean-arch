package postgres

import "database/sql"

type ExpenseRepository struct {
	Conn *sql.DB
}

func NewExpenseRepository(conn *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{Conn: conn}
}
