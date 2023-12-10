package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bar struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Image   string             `json:"image,omitempty" bson:"image,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
}

func GetBar(c *fiber.Ctx) error {
	bar := db.Bares

	busqueda, err := bar.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}
	defer busqueda.Close(context.Background())

	var bars []bson.M
	if err = busqueda.All(context.Background(), &bars); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}
	return c.Status(fiber.StatusAccepted).JSON(bars)
}

func GetBarById(c *fiber.Ctx) error {
	bar := db.Bares

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}

	busqueda := bar.FindOne(context.Background(), bson.M{"_id": objID})

	var barres Bar
	if err = busqueda.Decode(&barres); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}
	return c.Status(fiber.StatusAccepted).JSON(barres)
}

func CreateBar(c *fiber.Ctx) error {
	bar := db.Bares

	data := new(Bar)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el bar"})
	}

	insertion, err := bar.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el bar"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el bar"})
	}else{
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Bar creado"})
	}
}

func UpdateBar(c *fiber.Ctx) error {
	bar := db.Bares

	data := new(Bar)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el bar"})
	}


	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el bar"})
	}

	_, err = bar.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el bar"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Bar actualizado correctamente"})
}

func DeleteBar(c *fiber.Ctx) error {
	bar := db.Bares

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}

	_, err = bar.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el bar"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Bar eliminado correctamente"})
}