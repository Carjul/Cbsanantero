package controllers

import (
	"context"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHospedaje(c *fiber.Ctx) error {
	hospedaje := db.Hospedaje

	busqueda, err := hospedaje.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}
	defer busqueda.Close(context.Background())

	var hospedajes []bson.M
	if err = busqueda.All(context.Background(), &hospedajes); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}
	return c.Status(fiber.StatusAccepted).JSON(hospedajes)
}

func GetHospedajeById(c *fiber.Ctx) error {
	hospedajes := db.Hospedaje

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}

	var hospedaje bson.M

	err = hospedajes.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&hospedaje)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}

	return c.Status(fiber.StatusAccepted).JSON(hospedaje)
}

func CreateHospedaje(c *fiber.Ctx) error {
	hospedaje := db.Hospedaje
	customer := db.Customer

	data := new(models.Hospedaje)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hospedaje"})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	customerID := form.Value["customer_id"]
	data.CustomerID = customerID[0]
	idc := data.CustomerID

	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El customer_id no es valido"})
	}

	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})

	var customerData models.Customer

	if err = busqueda.Decode(&customerData); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})

	}

	data.Status = "Activo"

	// Obtiene los archivos subidos
	files := form.File["image"]

	ImageFile := files[0]
	if ImageFile == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})
	}
	UrlCloudinary := config.UploadImage(ImageFile)
	data.Image = UrlCloudinary

	insertion, err := hospedaje.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hospedaje"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el hospedaje"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hospedaje creado correctamente"})
	}
}

func UpdateHospedaje(c *fiber.Ctx) error {
	hospedaje := db.Hospedaje

	data := new(models.Hospedaje)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos del hospedaje no son corectos"})
	}

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	files := form.File["image"]

	ImageFile := files[0]
	if ImageFile != nil {
		UrlCloudinary := config.UploadImage(ImageFile)
		data.Image = UrlCloudinary
	}

	update := bson.M{
		"$set": data,
	}

	_, err = hospedaje.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el hospedaje"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hospedaje actualizado correctamente"})
}

func DeleteHospedaje(c *fiber.Ctx) error {
	hospedajes := db.Hospedaje

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}
	var hospedaje bson.M

	err = hospedajes.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&hospedaje)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}

	_, err = hospedajes.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el hospedaje"})
	}
	config.DeleteImage(hospedaje["image"].(string))

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hospedaje eliminado correctamente"})
}
