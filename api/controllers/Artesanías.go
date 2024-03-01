package controllers

import (
	"github.com/cbsanantero/services"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artesanias struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Image   string             `json:"image,omitempty" bson:"image,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

func GetArtesanias(c *fiber.Ctx) error {
	artesanias:= services.GetArtesanias()
	return c.Status(fiber.StatusAccepted).JSON(artesanias)
	
}

func GetArtesaniasById(c *fiber.Ctx) error {
	id := c.Params("id")
	artesania:= services.GetArtesaniasById(id)
	if artesania == "error al buscar artesania"{
		return c.Status(fiber.StatusNotFound).JSON(Message{Msg: "Error al buscar artesanias"})
	
	}
	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func CreateArtesanias(c *fiber.Ctx) error {
	data := new(Artesanias)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}
	 argumen := Artesanias{
		Name: data.Name,
		Address: data.Address,
		Image: data.Image,
		Phone: data.Phone,
		Description: data.Description,
		Status: "Activo",
		CustomerID: data.CustomerID,
	} 
	artesania := services.CreateArtesanias((*services.Artesanias)(&argumen))
 
	return c.Status(fiber.StatusCreated).JSON(artesania)
}

func UpdateArtesanias(c *fiber.Ctx) error {
	
	data := new(Artesanias)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}

	id := c.Params("id")

	argument := Artesanias{
		Name: data.Name,
		Address: data.Address,
		Image: data.Image,
		Phone: data.Phone,
		Description: data.Description,
		Status: data.Status,
		CustomerID: data.CustomerID,

	}
	artesania := services.UpdateArtesanias( (*services.Artesanias)(&argument),id)
	 
	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func DeleteArtesanias(c *fiber.Ctx) error {
	
	id := c.Params("id")

	artesania := services.DeleteArtesanias(id)

	return c.Status(fiber.StatusAccepted).JSON(artesania)
}
