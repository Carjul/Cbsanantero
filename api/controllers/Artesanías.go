package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	artesanias := db.Artesanias

	busqueda, err := artesanias.Find(context.TODO(), bson.M{"status":"Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}
	defer busqueda.Close(context.Background())

	var artesania []bson.M
	if err = busqueda.All(context.Background(), &artesania); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}
	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func GetArtesaniasById(c *fiber.Ctx) error {
	artesanias := db.Artesanias

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}

	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID,"status":"Activo"})

	var artesania Artesanias
	if err = busqueda.Decode(&artesania); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}
	return c.Status(fiber.StatusAccepted).JSON(artesania)
}

func CreateArtesanias(c *fiber.Ctx) error {
	artesanias := db.Artesanias
	customer := db.Customer

	data := new(Artesanias)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}

	idc:= data.CustomerID

	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El _id no es valido"})
	}

	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})

	var customerData Customer

	if err = busqueda.Decode(&customerData); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})
	
	}
	data.Status = "Activo"
	

	result, err := artesanias.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar la artesania"})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear la artesania"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Artesania creada"})
	}

}

func UpdateArtesanias(c *fiber.Ctx) error {
	artesanias := db.Artesanias

	data := new(Artesanias)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo decodificar la artesania"})
	}

	objID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}

	result, err := artesanias.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar la artesania"})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar la artesania"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Artesania actualizada"})
	}
}

func DeleteArtesanias(c *fiber.Ctx) error {
	artesanias := db.Artesanias

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la artesania"})
	}

	result, err := artesanias.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar la artesania"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar la artesania"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Artesania eliminada"})
	}
}