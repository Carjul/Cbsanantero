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
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Image    string             `json:"image,omitempty" bson:"image,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Rol      string             `json:"rol,omitempty" bson:"rol,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
}

func GetCustumer(c *fiber.Ctx) error {
	customer := db.Customer

	busqueda, err := customer.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer busqueda.Close(context.Background())

	var customers []bson.M

	if err = busqueda.All(context.Background(), &customers); err != nil {
		log.Println(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(customers)
}

func GetCustumerById(c *fiber.Ctx) error {
	customers := db.Customer

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	var customer bson.M

	err = customers.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})
	}

	return c.Status(fiber.StatusAccepted).JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customers := db.Customer

	customer := new(Customer)

	if err := c.BodyParser(customer); err != nil {
		log.Println(err)
	}
	customer.Status = "Inactivo"
	customer.Rol = "Cliente"
	result, err := customers.InsertOne(context.Background(), customer)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar el cliente"})
	}
	if result.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar el cliente"})
	} else {
		return c.Status(fiber.StatusCreated).JSON(Message{Msg: "Cliente insertado"})
	}

}

func UpdateCustomer(c *fiber.Ctx) error {
	customers := db.Customer

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	customer := new(Customer)

	if err := c.BodyParser(customer); err != nil {
		log.Println(err)
	}

	update := bson.M{
		"$set": customer,
	}

	result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Println(err)
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el cliente"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Cliente actualizado"})
	}
}

func DeleteCustomer(c *fiber.Ctx) error {
	customers := db.Customer

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	result, err := customers.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
	}
	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el cliente"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Cliente eliminado"})
	}
}

func UpdateCustomerStatus (c *fiber.Ctx) error {
    type Status struct {
        Status string `json:"status,omitempty" bson:"status,omitempty"`
    }

    customers := db.Customer

    id := c.Params("id")

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var customerStatus Status

    if err := c.BodyParser(&customerStatus); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    x := Status{Status: customerStatus.Status}
    update := bson.M{"$set":x}

    result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update customer"})
    }

    if result.ModifiedCount == 0 {
        return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "No se pudo actualizar el cliente"})
    } else {
        return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Cliente actualizado"})
    }
}

func UpdateCustomerRol (c *fiber.Ctx) error {
    type Rol struct {
		Rol string `json:"rol,omitempty" bson:"rol,omitempty"`
	}

    customers := db.Customer

    id := c.Params("id")

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var customerRol Rol

    if err := c.BodyParser(&customerRol); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    x := Rol{Rol: customerRol.Rol}
    update := bson.M{"$set":x}

    result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update rol"})
    }

    if result.ModifiedCount == 0 {
        return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "No se pudo actualizar el cliente"})
    } else {
        return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Cliente actualizado"})
    }
}