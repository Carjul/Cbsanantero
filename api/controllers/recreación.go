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

func GetRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	busqueda, err := recreacion.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}
	defer busqueda.Close(context.Background())

	var recreaciones []bson.M
	if err = busqueda.All(context.Background(), &recreaciones); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}
	return c.Status(fiber.StatusAccepted).JSON(recreaciones)
}

func GetRecreacionById(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	var recreaciones bson.M

	err = recreacion.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&recreaciones)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	return c.Status(fiber.StatusAccepted).JSON(recreaciones)
}

func CreateRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion
	customer := db.Customer

	data := new(models.Recreacion)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos insertados estan malos"})
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
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El _id no es valido"})
	}

	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})

	var customerData models.Customer

	if err = busqueda.Decode(&customerData); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})

	}
	data.Status = "Activo"

	files := form.File["image"]

	ImageFile := files[0]
	if ImageFile == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})
	}
	UrlCloudinary := config.UploadImage(ImageFile)
	data.Image = UrlCloudinary

	insertion, err := recreacion.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar la recreacion"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Recreacion creada correctamente"})
	}
}

func UpdateRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	id := c.Params("id")

	data := new(models.Recreacion)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos de la recreacion no son corectos"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	files := form.File["image"]

	ImageFile := files[0]
	if ImageFile != nil {
		UrlCloudinary := config.UploadImage(ImageFile)
		data.Image = UrlCloudinary
	}

	_, err = recreacion.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar la recreacion"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Recreacion actualizada correctamente"})
}

func DeleteRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	var recreaciones bson.M

	err = recreacion.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&recreaciones)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	_, err = recreacion.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar la recreacion"})
	}
	config.DeleteImage(recreaciones["image"].(string))
	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Recreacion eliminada correctamente"})
}
