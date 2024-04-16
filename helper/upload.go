package helper

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

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

func UploadDoc(file *multipart.FileHeader, id string) (string, error) {

	src, err := file.Open()
	if err != nil {
		log.Println("error membuka file", err.Error())
		return "", err
	}
	defer src.Close()
	dir := "uploads"
	dirUser := dir + "/" + id
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println(dir, "does not exist")
		os.Mkdir(dir, 0644)

	} else {
		log.Println("The provided directory named", dir, "exists")
	}
	if _, err := os.Stat(dirUser); os.IsNotExist(err) {
		log.Println(dirUser, "does not exist")
		os.Mkdir(dirUser, 0644)

	} else {
		log.Println("The provided directory named", dir, "exists")
	}
	filename := strings.Replace(file.Filename, " ", "-", -1)
	// Destination
	dst, err := os.Create(dirUser + "/" + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return ("https://inbuscap-server.my.id/" + dirUser + "/" + filename), nil
}
