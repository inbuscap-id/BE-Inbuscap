package main

import (
	"BE-Inbuscap/config"
	invest_data "BE-Inbuscap/features/invest/data"
	invest_handler "BE-Inbuscap/features/invest/handler"
	invest_services "BE-Inbuscap/features/invest/services"
	proposal_data "BE-Inbuscap/features/proposal/data"
	proposal_handler "BE-Inbuscap/features/proposal/handler"
	proposal_services "BE-Inbuscap/features/proposal/services"
	transaction_data "BE-Inbuscap/features/transaction/data"
	transaction_handler "BE-Inbuscap/features/transaction/handler"
	transaction_services "BE-Inbuscap/features/transaction/services"
	user_data "BE-Inbuscap/features/user/data"
	user_handler "BE-Inbuscap/features/user/handler"
	user_services "BE-Inbuscap/features/user/services"
	"BE-Inbuscap/routes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)
	config.Migrate(db, &user_data.User{}, &user_data.Proposal{}, &user_data.Investment{}, &user_data.Report{}, &transaction_data.Transaction{})

	userData := user_data.New(db)
	userService := user_services.NewService(userData)
	userHandler := user_handler.NewHandler(userService)

	proposalData := proposal_data.New(db)
	proposalService := proposal_services.Service(proposalData)
	proposalHandler := proposal_handler.NewHandler(proposalService)

	investData := invest_data.New(db)
	investService := invest_services.Service(investData)
	investHandler := invest_handler.NewHandler(investService)

	transactionData := transaction_data.New(db)
	transactionService := transaction_services.New(transactionData)
	transactionHandler := transaction_handler.New(transactionService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, userHandler, proposalHandler, investHandler, transactionHandler)
	e.Logger.SetLevel(log.INFO)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
