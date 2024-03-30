package handler

import (
	"BE-Inbuscap/features/proposal"
	"BE-Inbuscap/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	s proposal.Services
}

func NewHandler(service proposal.Services) proposal.Controller {
	return &controller{
		s: service,
	}
}

func (ct *controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) Archive() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}