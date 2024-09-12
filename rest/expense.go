package rest

import (
	"net/http"

	"github.com/brown-kaew/go-try-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type ExpenseService interface {
	Create(expense domain.Expense) error
}

type ExpenseHandler struct {
	Service ExpenseService
}

func NewExpenseHandler(e *echo.Echo, svc ExpenseService) {
	handler := &ExpenseHandler{
		Service: svc,
	}
	e.POST("/expenses", handler.createNewExpenseHandler)

}

func (h *ExpenseHandler) createNewExpenseHandler(c echo.Context) error {
	var expense domain.Expense
	err := c.Bind(&expense)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.Service.Create(expense)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, expense)

}
