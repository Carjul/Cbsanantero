package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trasporte struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tipo    string             `json:"tipo,omitempty" bson:"tipo,omitempty"`
	Image  string             `json:"image,omitempty" bson:"image,omitempty"`
	Placa   string             `json:"placa,omitempty" bson:"placa,omitempty"`
	Conductor   string         `json:"conductor,omitempty" bson:"conductor,omitempty"`
	Celular   string           `json:"celular,omitempty" bson:"celular,omitempty"`
}

func GetTrasporte(c *fiber.Ctx) error {
	trasporte := db.Traporte

	busqueda, err := trasporte.Find(context.TODO(), bson.M{})
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
	trasportes := db.Traporte

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

	return c.Status(fiber.StatusAccepted).JSON(trasporte)
}

func CreateTrasporte(c *fiber.Ctx) error {
	trasporte := db.Traporte

	data := new(Trasporte)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos del trasporte no son corectos"})
	}

	insertion, err := trasporte.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el trasporte"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el trasporte"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte creado correctamente"})
	}
}

func UpdateTrasporte(c *fiber.Ctx) error {
	trasporte := db.Traporte

	data := new(Trasporte)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Los datos del trasporte no son corectos"})
	}

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	_, err = trasporte.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el trasporte"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte actualizado correctamente"})
}

func DeleteTrasporte(c *fiber.Ctx) error {
	trasporte := db.Traporte

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el trasporte"})
	}

	_, err = trasporte.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el trasporte"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Trasporte eliminado correctamente"})
}