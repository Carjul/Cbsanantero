package controllers

import (
	"github.com/cbsanantero/config"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Msg string `json:"message"`
}

func Init(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(Message{
		Msg: "Bienvenido a la API de Costa Brisa",
	})
}
func Upload(c *fiber.Ctx) error {
	// Parsea el formulario multipart
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	files := form.File["image"]
	/* for _, file := range files {
		// Guarda el archivo en el sistema de archivos
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename)); err != nil {
			return err
		}
	} */
	x := files[0]
	y := config.UploadImage2(x)
	return c.JSON(y)
}

type Imagen struct {
	Url string `json:"image"`
}
