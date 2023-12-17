package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recreacion struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Services string             `json:"services,omitempty" bson:"services,omitempty"`
	Image   string             `json:"image,omitempty" bson:"image,omitempty"`
	Price  string             `json:"price,omitempty" bson:"price,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

func GetRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	busqueda, err := recreacion.Find(context.TODO(), bson.M{"status":"Activo"})
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

	err = recreacion.FindOne(context.TODO(), bson.M{"_id": objID,"status":"Activo"}).Decode(&recreaciones)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}

	return c.Status(fiber.StatusAccepted).JSON(recreaciones)
}

func CreateRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion
	customer := db.Customer

	data := new(Recreacion)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos insertados estan malos"})
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

	insertion, err := recreacion.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar la recreacion"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Recreacion creada correctamente"})
	}
}

func UpdateRecreacion(c *fiber.Ctx) error {
	recreacion := db.Recreacion

	data := new(Recreacion)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos de la recreacion no son corectos"})
	}

  

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar la recreacion"})
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

	_, err = recreacion.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar la recreacion"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Recreacion eliminada correctamente"})
}