package controllers

import (
	"context"
	"sync"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTrasporte(c *fiber.Ctx) error {
	trasporte := Instance.Database.Collection("Trasporte")

	busqueda, err := trasporte.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}
	defer busqueda.Close(context.Background())

	var trasportes []bson.M
	if err = busqueda.All(context.Background(), &trasportes); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}
	return c.Status(fiber.StatusAccepted).JSON(trasportes)
}

func GetTrasporteById(c *fiber.Ctx) error {
	trasportes := Instance.Database.Collection("Trasporte")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	var trasporte bson.M

	err = trasportes.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&trasporte)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	return c.Status(fiber.StatusAccepted).JSON(trasporte)
}

func CreateTrasporte(c *fiber.Ctx) error {
	trasporte := Instance.Database.Collection("Trasporte")
	customer := Instance.Database.Collection("Customer")

	data := new(models.Trasporte)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos del trasporte no son corectos"})
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
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "La imagen es requerida",
		})
	}
	ImageFile := files[0]
	var UrlCloudinary string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		UrlCloudinary = config.UploadImage(ImageFile)
		wg.Done()
	}()
	wg.Wait()

	if UrlCloudinary == "error al subir la imagen a cloudinary" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al subir la imagen",
		})
	}

	data.Image = UrlCloudinary

	insertion, err := trasporte.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el trasporte"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el trasporte"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte creado correctamente"})
	}
}

func UpdateTrasporte(c *fiber.Ctx) error {
	trasporte := Instance.Database.Collection("Trasporte")

	id := c.Params("id")

	data := new(models.Trasporte)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos del trasporte no son corectos"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	// Obtiene los archivos subidos
	files := form.File["image"]
	ImageFile := files[0]

	if ImageFile != nil {
		UrlCloudinary := config.UploadImage(ImageFile)
		data.Image = UrlCloudinary
	}

	_, err = trasporte.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el trasporte"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte actualizado correctamente"})
}

func DeleteTrasporte(c *fiber.Ctx) error {
	trasportes := Instance.Database.Collection("Trasporte")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	var trasporte bson.M

	err = trasportes.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&trasporte)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	_, err = trasportes.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el trasporte"})
	}
	config.DeleteImage(trasporte["image"].(string))
	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte eliminado correctamente"})
}
