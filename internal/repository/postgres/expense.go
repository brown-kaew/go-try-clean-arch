package postgres

import (
	"database/sql"

	"github.com/brown-kaew/go-try-clean-arch/domain"
	"github.com/lib/pq"
)

type ExpenseRepository struct {
	Conn *sql.DB
}

func NewExpenseRepository(conn *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{Conn: conn}
}

func (e *ExpenseRepository) Create(expense *domain.Expense) error {
	sql := `
	INSERT INTO expenses (title, amount, note, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	row := e.Conn.QueryRow(sql, expense.Title, expense.Amount, expense.Note, pq.Array(&expense.Tags))
	if err := row.Scan(&expense.Id); err != nil {
		return err
	}
	return nil
}
