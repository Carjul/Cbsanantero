package controllers

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Galeria struct {
	Photos      []map[string]interface{} `json:"photos,omitempty" bson:"photos,omitempty"`
	Description string                   `json:"description,omitempty" bson:"description,omitempty"`
	NegocioID   string                   `json:"negocio_id,omitempty" bson:"negocio_id,omitempty"`
	Status      string                   `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string                   `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

func GetGaleria(c *fiber.Ctx) error {
	galeria := Instance.Database.Collection("Galeria")
	id := c.Params("id")

	busqueda, err := galeria.Find(context.TODO(), bson.M{"negocio_id": id, "status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la galería"})
	}
	defer busqueda.Close(context.Background())

	var galerias []models.Galeria
	if err := busqueda.All(context.Background(), &galerias); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la galería"})
	}

	var result Galeria
	if len(galerias) != 0 {
		// Asigna las fotos y la descripción
		for _, galeria := range galerias {
			for _, photo := range galeria.Photos {
				obj := map[string]interface{}{"image": photo, "id": galeria.ID}
				result.Photos = append(result.Photos, obj)
			}
			result.Description = galeria.Description
			result.Status = galeria.Status
			result.CustomerID = galeria.CustomerID
			result.NegocioID = galeria.NegocioID
		}
	}

	return c.Status(fiber.StatusAccepted).JSON(result)
}

func CreateGaleria(c *fiber.Ctx) error {
	galeria := Instance.Database.Collection("Galeria")
	customer := Instance.Database.Collection("Customer")

	data := new(models.Galeria)

	// Parsea el formulario
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Obtiene los archivos subidos
	longitud := form.Value["logitud"]
	customerID := form.Value["customer_id"]
	negocioID := form.Value["negocio_id"]
	description := form.Value["description"] // Obtén la descripción del formulario

	num, err := strconv.Atoi(longitud[0])

	if err != nil {
		fmt.Println("Error al convertir la cadena a un número:", err)
	}
	for i := 0; i < num; i++ {
		str := strconv.Itoa(i)
		name := "image" + str
		files := form.File[name]
		ImageFile := files[0]

		if ImageFile.Filename == "" {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})
		}

		var UrlCloudinary string
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			UrlCloudinary = config.UploadImage(ImageFile)
			data.Photos = append(data.Photos, UrlCloudinary)
			wg.Done()

		}()
		wg.Wait()
	}

	data.CustomerID = customerID[0]
	data.NegocioID = negocioID[0]
	data.Status = "Activo"
	data.Description = description[0] // Asigna la descripción al modelo

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

	if data.Photos == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})

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

func UpdateGaleria(c *fiber.Ctx) error {
	galeria := Instance.Database.Collection("Galeria")
	type UrlImage struct {
		Url string `json:"photo"`
	}
	id := c.Params("id")

	data := new(UrlImage)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la imagen"})
	}
	fmt.Println(data.Url)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El id no es valido"})
	}
	var galeriaData models.Galeria
	galeria.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&galeriaData)

	if galeriaData.Photos == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la photo"})

	}

	result := eliminarElemento(galeriaData.Photos, data.Url)

	_, err = galeria.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"photos": result}})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar la galeria"})
	}

	config.DeleteImage(data.Url)

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "galeria elinimada"})

}

func eliminarElemento(arr []string, elemento string) []string {
	// Buscar el índice del elemento a eliminar
	indice := -1
	for i, v := range arr {
		if v == elemento {
			indice = i
			break
		}
	}

	// Si el elemento existe, eliminarlo
	if indice != -1 {
		// Reconstruir el array sin el elemento
		arr = append(arr[:indice], arr[indice+1:]...)
	}

	return arr
}
