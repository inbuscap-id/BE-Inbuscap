package main

import (
	"BE-Inbuscap/config"
	invest_data "BE-Inbuscap/features/invest/data"
	invest_handler "BE-Inbuscap/features/invest/handler"
	invest_services "BE-Inbuscap/features/invest/services"
	post_data "BE-Inbuscap/features/post/data"
	post_handler "BE-Inbuscap/features/post/handler"
	post_services "BE-Inbuscap/features/post/services"
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
	config.Migrate(db, &user_data.User{}, &post_data.Post{}, &invest_data.Invest{})

	userData := user_data.New(db)
	userService := user_services.NewService(userData)
	userHandler := user_handler.NewHandler(userService)

	postData := post_data.New(db)
	postService := post_services.Service(postData)
	postHandler := post_handler.NewHandler(postService)

	investData := invest_data.New(db)
	investService := invest_services.Service(investData)
	investHandler := invest_handler.NewHandler(investService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, userHandler, postHandler, investHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
