package middlewares

import (
	"net/http"

	"BE-Inbuscap/helper"

	golangjwt "github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole, _ := DecodeRole(c.Get("user").(*golangjwt.Token))

		if !userRole {
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, helper.ErrorAuthorization))
		}
		return next(c)
	}
}

func CheckStatus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, _ := DecodeStatus(c.Get("user").(*golangjwt.Token))

		if status != 1 {
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, helper.ErrorAccountActivation))
		}

		return next(c)
	}
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\nhost=${host}, uri=${uri}, user_agent=${user_agent}, time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
}
