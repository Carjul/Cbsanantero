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
	/* go config.UploadImageLocal(data.Image)
	time.Sleep(1 * time.Second)
	data.Image = config.UploadImage() */
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	files := form.File["image"]
	customerID := form.Value["customer_id"]

	x := files[0]
	y := config.UploadImage2(x)
	data.Image = y
	data.CustomerID = customerID[0]

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
	customerID := form.Value["customer_id"]

	x := files[0]
	y := config.UploadImage2(x)
	data.Image = y
	data.CustomerID = customerID[0]
	/* if data.Image != "" {
		go config.UploadImageLocal(data.Image)
		time.Sleep(1 * time.Second)
		data.Image = config.UploadImage()
	} */
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
	hospedaje := db.Hospedaje

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el hospedaje"})
	}

	_, err = hospedaje.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el hospedaje"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Hospedaje eliminado correctamente"})
}
