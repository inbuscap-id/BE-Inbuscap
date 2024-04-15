package report

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Create() echo.HandlerFunc
	Edit() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Archive() echo.HandlerFunc
}

type Model interface {
	Create() error
	Edit() error
	GetAll() error
	GetDetail() error
	Delete() error
	Archive() error
}

type Services interface {
	Create() error
	Edit() error
	GetAll() error
	GetDetail() error
	Delete() error
	Archive() error
}
