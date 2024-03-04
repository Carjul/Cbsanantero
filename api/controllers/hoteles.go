package controllers

import (
	"context"
	"time"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHoteles(c *fiber.Ctx) error {
	hoteles := db.Hoteles

	busqueda, err := hoteles.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hotel"})
	}
	defer busqueda.Close(context.Background())

	var hoteless []bson.M
	if err = busqueda.All(context.Background(), &hoteless); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hotel"})
	}
	return c.Status(fiber.StatusAccepted).JSON(hoteless)
}

func GetHotelesById(c *fiber.Ctx) error {
	hoteles := db.Hoteles

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hotel"})
	}

	var hotel bson.M

	err = hoteles.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&hotel)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hotel"})
	}

	return c.Status(fiber.StatusAccepted).JSON(hotel)
}

func CreateHoteles(c *fiber.Ctx) error {
	hoteles := db.Hoteles
	customer := db.Customer

	data := new(models.Hoteles)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hotel"})
	}
	idc := data.CustomerID

	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El _id no es valido"})
	}

	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})

	var customerData models.Customer

	if err = busqueda.Decode(&customerData); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})

	}
	data.Status = "Activo"
	go config.UploadImageLocal(data.Image)
	time.Sleep(1 * time.Second)
	data.Image = config.UploadImage()

	insertion, err := hoteles.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hotel"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hotel"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hotel creado correctamente"})
	}
}

func UpdateHoteles(c *fiber.Ctx) error {
	hoteles := db.Hoteles

	data := new(models.Hoteles)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el hotel"})
	}
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el hotel"})
	}
	if data.Image != "" {
		go config.UploadImageLocal(data.Image)
		time.Sleep(1 * time.Second)
		data.Image = config.UploadImage()
	}

	_, err = hoteles.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el hotel"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hotel actualizado correctamente"})

}

func DeleteHoteles(c *fiber.Ctx) error {
	hoteles := db.Hoteles

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el hotel"})
	}

	_, err = hoteles.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el hotel"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hotel eliminado correctamente"})
}
