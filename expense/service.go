package expense

import "github.com/brown-kaew/go-try-clean-arch/domain"

type ExpenseRepository interface {
	Create(expense *domain.Expense) error
	GetById(id int) (*domain.Expense, error)
}

type Service struct {
	expenseRepo ExpenseRepository
}

func NewService(e ExpenseRepository) *Service {
	return &Service{expenseRepo: e}
}

func (s *Service) Create(expense *domain.Expense) error {
	return s.expenseRepo.Create(expense)
}

func (s *Service) GetById(id int) (*domain.Expense, error) {
	return s.expenseRepo.GetById(id)
}
