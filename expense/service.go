package expense

type ExpenseRepository interface {
}

type Service struct {
	expenseRepo ExpenseRepository
}

func New(e ExpenseRepository) *Service {
	return &Service{expenseRepo: e}
}
