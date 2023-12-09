package controllers

import (
	"context"
	"log"

	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
}

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

func GetCustumerById(c *fiber.Ctx) error {
    customers := db.Client.Database("test").Collection("Customer")

    id := c.Params("id")

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        panic(err)
    }

    var customer bson.M

    err = customers.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
    if err != nil {
        panic(err)
    }

    return c.Status(fiber.StatusAccepted).JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customers := db.Client.Database("test").Collection("Customer")

	customer := new(Customer)

	if err := c.BodyParser(customer); err != nil {
		panic(err)
	}

	result, err := customers.InsertOne(context.Background(), customer)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar el cliente"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

