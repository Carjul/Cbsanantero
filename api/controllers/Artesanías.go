package controllers

import (
	"fmt"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db/models"
	"github.com/cbsanantero/services"
	"github.com/gofiber/fiber/v2"
)

func GetArtesanias(c *fiber.Ctx) error {
	artesanias := services.GetArtesanias()
	return c.Status(fiber.StatusAccepted).JSON(artesanias)

}

func GetArtesaniasById(c *fiber.Ctx) error {
	id := c.Params("id")
	artesania := services.GetArtesaniasById(id)
	if artesania == "error al buscar artesania" {
		return c.Status(fiber.StatusNotFound).JSON(Message{Msg: "Error al buscar artesanias"})

	}
	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func CreateArtesanias(c *fiber.Ctx) error {
	data := new(models.Artesanias)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	files := form.File["image"]

	x := files[0]
	y := config.UploadImage2(x)
	fmt.Println(data)
	argumen := models.Artesanias{
		Name:        data.Name,
		Address:     data.Address,
		Image:       y,
		Phone:       data.Phone,
		Description: data.Description,
		Status:      "Activo",
		CustomerID:  data.CustomerID,
	}
	fmt.Println(argumen)
	artesania := services.CreateArtesanias((*models.Artesanias)(&argumen))

	return c.Status(fiber.StatusCreated).JSON(artesania)
}

func UpdateArtesanias(c *fiber.Ctx) error {

	data := new(models.Artesanias)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}

	id := c.Params("id")

	argument := models.Artesanias{
		Name:        data.Name,
		Address:     data.Address,
		Image:       data.Image,
		Phone:       data.Phone,
		Description: data.Description,
		Status:      data.Status,
		CustomerID:  data.CustomerID,
	}
	artesania := services.UpdateArtesanias((*models.Artesanias)(&argument), id)

	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func DeleteArtesanias(c *fiber.Ctx) error {

	id := c.Params("id")

	artesania := services.DeleteArtesanias(id)

	return c.Status(fiber.StatusAccepted).JSON(artesania)
}
