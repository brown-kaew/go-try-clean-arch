package rest

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/brown-kaew/go-try-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type ExpenseService interface {
	Create(expense *domain.Expense) error
	GetById(id int) (*domain.Expense, error)
	FetchAll() ([]domain.Expense, error)
}

type ExpenseHandler struct {
	Service ExpenseService
}

func NewExpenseHandler(e *echo.Echo, svc ExpenseService) {
	handler := &ExpenseHandler{
		Service: svc,
	}
	e.POST("/expenses", handler.Store)
	e.GET("/expenses/:id", handler.GetById)
	e.GET("/expenses", handler.FetchAll)
}

func (h *ExpenseHandler) Store(c echo.Context) error {
	var expense domain.Expense
	if err := c.Bind(&expense); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := h.Service.Create(&expense); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	expense, err := h.Service.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) FetchAll(c echo.Context) error {
	expenses, err := h.Service.FetchAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, expenses)
}
