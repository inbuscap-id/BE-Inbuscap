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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Logger.Fatal(e.Start(":1323"))
}
