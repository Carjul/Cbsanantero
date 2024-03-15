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

func GetTour(c *fiber.Ctx) error {
	tour := Instance.Database.Collection("Tour")

	busqueda, err := tour.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}
	defer busqueda.Close(context.Background())

	var tours []bson.M
	if err = busqueda.All(context.Background(), &tours); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}
	return c.Status(fiber.StatusAccepted).JSON(tours)
}

func GetTourById(c *fiber.Ctx) error {
	tours := Instance.Database.Collection("Tour")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	var tour bson.M

	err = tours.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&tour)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	return c.Status(fiber.StatusAccepted).JSON(tour)
}

func CreateTour(c *fiber.Ctx) error {
	tour := Instance.Database.Collection("Tour")
	customer := Instance.Database.Collection("Customer")

	data := new(models.Tour)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
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

	// Obtiene los archivos subidos
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

	result, err := tour.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour creado correctamente"})
	}
}

func UpdateTour(c *fiber.Ctx) error {
	tour := Instance.Database.Collection("Tour")

	id := c.Params("id")

	data := new(models.Tour)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
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

	result, err := tour.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour actualizado correctamente"})
	}
}

func DeleteTour(c *fiber.Ctx) error {
	tours := Instance.Database.Collection("Tour")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	var tour bson.M

	err = tours.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&tour)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	result, err := tours.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el tour"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el tour"})
	} else {
		config.DeleteImage(tour["image"].(string))
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour eliminado correctamente"})
	}
}
