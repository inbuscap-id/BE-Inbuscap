package utils

import (
	"BE-Inbuscap/config"
	"context"
	"fmt"
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

	env := config.InitConfig()
	cld, err := cloudinary.NewFromURL(fmt.Sprintf("cloudinary://%s:%s@%s", env.Cloudinary_API_Key, env.Cloudinary_API_Secret, env.Cloudinary_API_CLD))
	if err != nil {
		log.Println("error connecting to cloudinary:", err.Error())
		return "", err
	}

	resp, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: env.Folder, ResourceType: "auto"})
	if err != nil {
		log.Println("cloudinary upload error:", err.Error())
		return "", err
	}

	return resp.SecureURL, nil
}
