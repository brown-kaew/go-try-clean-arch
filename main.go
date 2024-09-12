package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brown-kaew/go-try-clean-arch/expense"
	"github.com/brown-kaew/go-try-clean-arch/internal/repository/postgres"
	"github.com/brown-kaew/go-try-clean-arch/rest"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	defaultAddress = ":2565"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	if err = dbConn.Ping(); err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	log.Println("db started")

	e := echo.New()

	// Prepare Repository
	expenseRepo := postgres.NewExpenseRepository(dbConn)

	// Build service Layer
	svc := expense.NewService(expenseRepo)
	rest.NewExpenseHandler(e, svc)

	go func() {
		if err := e.Start(defaultAddress); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Info("Server stopped")
}
