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

func (e *ExpenseRepository) GetById(id int) (*domain.Expense, error) {
	stmt, err := e.Conn.Prepare("SELECT * FROM expenses WHERE id=$1")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	var expense domain.Expense
	err = row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (e *ExpenseRepository) FetchAll() ([]domain.Expense, error) {
	stmt, err := e.Conn.Prepare("SELECT * FROM expenses")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var expenses []domain.Expense
	for rows.Next() {
		var expense domain.Expense
		err := rows.Scan(
			&expense.Id,
			&expense.Title,
			&expense.Amount,
			&expense.Note,
			pq.Array(&expense.Tags),
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
