package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Image  string             `json:"image,omitempty" bson:"image,omitempty"`
	Trasporte   string             `json:"trasporte,omitempty" bson:"trasporte,omitempty"`
	Price   string             `json:"price,omitempty" bson:"price,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

func GetTour(c *fiber.Ctx) error {
	tour := db.Tour

	busqueda, err := tour.Find(context.TODO(), bson.M{"status":"Activo"})
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
	tours := db.Tour

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	var tour bson.M

	err = tours.FindOne(context.TODO(), bson.M{"_id": objID,"status":"Activo"}).Decode(&tour)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	return c.Status(fiber.StatusAccepted).JSON(tour)
}

func CreateTour(c *fiber.Ctx) error {
	tour := db.Tour
	customer := db.Customer

	data := new(Tour)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
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

	result, err := tour.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el tour"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour creado correctamente"})
	}
}

func UpdateTour(c *fiber.Ctx) error {
	tour := db.Tour

	data := new(Tour)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
	}

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el tour"})
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
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour actualizado correctamente"})
	}
}

func DeleteTour(c *fiber.Ctx) error {
	tour := db.Tour

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el tour"})
	}

	result, err := tour.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el tour"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el tour"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Tour eliminado correctamente"})
	}
}