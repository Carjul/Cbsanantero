package controllers

import (
	"sync"
	"context"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gofiber/fiber/v2"
)



func GetArtesania() interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	busqueda, err := artesanias.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return "error al buscar artesania"
	}
	defer busqueda.Close(context.Background())

	var artesania []bson.M
	if err = busqueda.All(context.Background(), &artesania); err != nil {
		return "No se pudo encontrar la artesania"
	}
	return artesania
}

func GetArtesaniaById(id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No existe el ID de la artesania"}
	}
	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID, "status": "Activo"})
	var artesania models.Artesanias
	if err := busqueda.Decode(&artesania); err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}
	return artesania
}

func CreateArtesania(data *models.Artesanias) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	customer := Instance.Database.Collection("Customer")
	idc := data.CustomerID
	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return Message{Msg: "El _id Customer no es valido"}
	}
	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})
	var customerData models.Customer
	if err = busqueda.Decode(&customerData); err != nil {
		return Message{Msg: "No se pudo encontrar el cliente"}
	}
	_, err = artesanias.InsertOne(context.Background(), data)
	if err != nil {
		return Message{Msg: "No se pudo crear la artesania"}
	}
	return Message{Msg: "Artesania creada"}

}

func UpdateArtesania(data *models.Artesanias, id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}

	result, err := artesanias.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return Message{Msg: "error al actualizar la artesania"}
	}
	if result.ModifiedCount == 0 {
		return Message{Msg: "No se pudo actualizar la artesania"}
	}
	return Message{Msg: "Artesania actualizada"}
}

func DeleteArtesania(id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "_id de la artesania no valido"}
	}
	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID})
	var artesania models.Artesanias
	if err := busqueda.Decode(&artesania); err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}

	result, err := artesanias.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return Message{Msg: "error al eliminar la artesania"}
	}
	if result.DeletedCount == 0 {
		return Message{Msg: "No se pudo eliminar la artesania"}
	}
	config.DeleteImage(artesania.Image)
	return Message{Msg: "Artesania eliminada"}
}



func GetArtesanias(c *fiber.Ctx) error {
	artesanias := GetArtesania()
	return c.Status(fiber.StatusAccepted).JSON(artesanias)

}

func GetArtesaniasById(c *fiber.Ctx) error {
	id := c.Params("id")
	artesania := GetArtesaniaById(id)
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
	customerID := form.Value["customer_id"]
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


	argumen := models.Artesanias{
		Name:        data.Name,
		Address:     data.Address,
		Image:       UrlCloudinary,
		Phone:       data.Phone,
		Description: data.Description,
		Status:      "Activo",
		CustomerID:  customerID[0],
	}

	artesania := CreateArtesania((*models.Artesanias)(&argumen))

	return c.Status(fiber.StatusCreated).JSON(artesania)
}

func UpdateArtesanias(c *fiber.Ctx) error {

	data := new(models.Artesanias)
	id := c.Params("id")

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	customerID := form.Value["customer_id"]
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


	argument := models.Artesanias{
		Name:        data.Name,
		Address:     data.Address,
		Image:       UrlCloudinary,
		Phone:       data.Phone,
		Description: data.Description,
		Status:      data.Status,
		CustomerID:  customerID[0],
	}
	artesania := UpdateArtesania((*models.Artesanias)(&argument), id)

	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func DeleteArtesanias(c *fiber.Ctx) error {

	id := c.Params("id")

	artesania := DeleteArtesania(id)

	return c.Status(fiber.StatusAccepted).JSON(artesania)
}
