package config

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryV2() *cloudinary.Cloudinary {
	//enviroment variables
	cloud_name := os.Getenv("CLOUD_NAME")
	api_key := os.Getenv("API_KEY")
	secret_key := os.Getenv("API_SECRET")
	if cloud_name == "" || api_key == "" || secret_key == "" {
		fmt.Println("No se encontraron las variables de entorno para cloudinary")
	} else {
		cld, _ := cloudinary.NewFromParams(cloud_name, api_key, secret_key)
		return cld
	}
	return nil
}

// Función que decodifica una imagen en base64 y la guarda en un archivo

func UploadImage2(image *multipart.FileHeader) string {

	cld := CloudinaryV2()

	// Configura los parámetros de la carga
	overwrite := true
	uploadParams := uploader.UploadParams{
		Folder:    "costa-brisa",
		Overwrite: &overwrite,
	}

	// Realiza la carga de la imagen a Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), image, uploadParams) // Change this line
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
		return "error al subir la imagen a cloudinary"
	}

	// Retorna la URL de la imagen cargada
	return uploadResult.SecureURL

}
