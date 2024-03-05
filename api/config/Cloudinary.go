package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var extension string

func CloudinaryV2() *cloudinary.Cloudinary {
	//enviroment variables
	cloud_name := "CLOUD_dm2q4hcus"
	api_key := "API_841516695972777"
	secret_key := "API__wQeZauUWPKrG8w92PegCzOqaik"
	if cloud_name == "" || api_key == "" || secret_key == "" {
		log.Fatal("No se encontraron las variables de entorno para cloudinary")
	} else {
		cld, _ := cloudinary.NewFromParams(cloud_name, api_key, secret_key)
		return cld
	}
	return nil
}

// Función que verifica si la cadena s contiene alguna de las palabras especificadas
func containsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// Función que decodifica una imagen en base64 y la guarda en un archivo
func UploadImageLocal(Url string) {

	if containsAny(Url, []string{"image/jpeg"}) {
		// Elimina el prefijo "data:image/jpeg;base64,"
		parts := strings.Split(Url, ",")
		if len(parts) != 2 {
			fmt.Println("Error: la cadena no está en un formato válido.")
		}

		// Decodifica la cadena base64
		data, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Println("Error decoding base64:", err)
		}

		// Escribe los datos decodificados en un archivo
		err = ioutil.WriteFile("./config/imagen.jpeg", data, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		}
		extension = ".jpeg"
	}

	if containsAny(Url, []string{"image/jpg"}) {
		// Elimina el prefijo "data:image/jpeg;base64,"
		parts := strings.Split(Url, ",")
		if len(parts) != 2 {
			fmt.Println("Error: la cadena no está en un formato válido.")
		}

		// Decodifica la cadena base64
		data, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Println("Error decoding base64:", err)
		}

		// Escribe los datos decodificados en un archivo
		err = ioutil.WriteFile("./config/imagen.jpg", data, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		}
		extension = ".jpg"
	}

	if containsAny(Url, []string{"image/png"}) {
		// Elimina el prefijo "data:image/jpeg;base64,"
		parts := strings.Split(Url, ",")
		if len(parts) != 2 {
			fmt.Println("Error: la cadena no está en un formato válido.")
		}

		// Decodifica la cadena base64
		data, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Println("Error decoding base64:", err)
		}

		// Escribe los datos decodificados en un archivo
		err = ioutil.WriteFile("./config/imagen.png", data, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		}
		extension = ".png"
	}

}

// Función que carga una imagen a Cloudinary
func UploadImage() string {

	cld := CloudinaryV2()

	// Configura los parámetros de la carga
	overwrite := true
	uploadParams := uploader.UploadParams{
		Folder:    "costa-brisa",
		Overwrite: &overwrite,
	}

	// Realiza la carga de la imagen a Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), "./config/imagen"+extension, uploadParams) // Change this line
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
		return "error al subir la imagen a cloudinary"
	}

	// Borra la imagen después de subirla a Cloudinary
	err = os.Remove("./config/imagen" + extension)
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
	extension = ""
	// Retorna la URL de la imagen cargada
	return uploadResult.SecureURL

}
