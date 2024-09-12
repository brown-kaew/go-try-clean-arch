package expense

import "github.com/brown-kaew/go-try-clean-arch/domain"

type ExpenseRepository interface {
	Create(expense domain.Expense) error
}

type Service struct {
	expenseRepo ExpenseRepository
}

func NewService(e ExpenseRepository) *Service {
	return &Service{expenseRepo: e}
}

func (s *Service) Create(expense domain.Expense) error {
	return nil
}
