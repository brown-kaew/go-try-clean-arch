package rest

import (
	"github.com/brown-kaew/go-try-clean-arch/expense"
	"github.com/labstack/echo/v4"
)

type ExpenseService interface {
}

func NewExpenseHandler(e *echo.Echo, svc ExpenseService) {
	handler := &expense.Service{}

	_ = handler
}
