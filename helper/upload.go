package helper

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

func SelectFile(c echo.Context, key string) (*multipart.FileHeader, error) {
	formHeader, err := c.FormFile(key)
	if err != nil {
		log.Println("error di helper:", err.Error())
		return nil, err
	}
	if formHeader == nil {
		log.Println("error di helper: file tidak terdeteksi")
		return nil, errors.New("file kosong")
	}
	return formHeader, nil
}
