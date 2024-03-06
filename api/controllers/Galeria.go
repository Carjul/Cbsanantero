package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGaleria(c *fiber.Ctx) error {
	galeria := db.Galeria
	id := c.Params("id")

	busqueda, err := galeria.Find(context.TODO(), bson.M{"negocio_id": id, "status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}
	defer busqueda.Close(context.Background())

	var galerias []bson.M
	if err = busqueda.All(context.Background(), &galerias); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}
	return c.Status(fiber.StatusAccepted).JSON(galerias)
}

func CreateGaleria(c *fiber.Ctx) error {
	galeria := db.Galeria
	customer := db.Customer

	data := new(models.Galeria)

	// Parsea el formulario
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	longitud := form.Value["logitud"]
	num, err := strconv.Atoi(longitud[0])

	if err != nil {
		fmt.Println("Error al convertir la cadena a un número:", err)
	}
	for i := 0; i < num; i++ {

		str := strconv.Itoa(i)
		name := "image" + str
		files := form.File[name]
		ImageFile := files[0]
		if ImageFile == nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})
		}
		UrlCloudinary := config.UploadImage(ImageFile)
		data.Photos = append(data.Photos, UrlCloudinary)
	}

	customerID := form.Value["customer_id"]
	negocioID := form.Value["negocio_id"]

	data.CustomerID = customerID[0]
	data.NegocioID = negocioID[0]
	data.Status = "Activo"

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

	insertion, err := galeria.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear la galeria"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear la galeria"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "galeria creado"})
	}
}

/* func UpdateGaleria(c *fiber.Ctx) error{

} */
