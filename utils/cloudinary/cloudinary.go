package config

import (
	"BE-Inbuscap/config"
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(image *multipart.FileHeader) (string, error) {
	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	url := config.InitConfig().CLOUDINARY_URL
	cld, err := cloudinary.NewFromURL(url)
	if err != nil {
		log.Println("error connecting to cloudinary:", err.Error())
		return "", err
	}

	resp, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		log.Println("cloudinary upload error:", err.Error())
		return "", err
	}

	imageUrl := resp.SecureURL

	return imageUrl, nil
}
