package controllers

import (
	"context"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCustumer(c *fiber.Ctx) error {
customer:= db.Client.Database("test").Collection("Customer")

busqueda, err := customer.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer busqueda.Close(context.Background())

var customers []bson.M

if err = busqueda.All(context.Background(), &customers); err != nil {
	panic(err)
}


	return c.Status(fiber.StatusAccepted).JSON(customers)
}
