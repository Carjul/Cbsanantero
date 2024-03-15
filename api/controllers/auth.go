package controllers

import (
	"context"
	"time"

	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	pasetoware "github.com/gofiber/contrib/paseto"
)

const secretSymmetricKey = "symmetric-secret-key (size = 32)"

func Login(c *fiber.Ctx) error {
	customer := Instance.Database.Collection("Customer")
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse body from JSON to struct
	var user Request
	err := c.BodyParser(&user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var userdb models.Customer
	// Check if user credentials are correct
	
	customer.FindOne(context.TODO(), bson.M{"email": user.Email, "password": user.Password}).Decode(&userdb)
	if userdb == (models.Customer{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(Message{Msg: "Credenciales incorrectas"})
	}

	// Create token and encrypt it
	encryptedToken, err := pasetoware.CreateToken([]byte(secretSymmetricKey), userdb.Name, 12*time.Hour, pasetoware.PurposeLocal)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	type Response struct {
		Token string          `json:"token"`
		User  models.Customer `json:"user"`
	}

	return c.JSON(Response{Token: encryptedToken, User: userdb})
}

func Restricted(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(string)
	return c.SendString("Welcome " + payload)
}

func LoginG(c *fiber.Ctx) error {
	Customer := Instance.Database.Collection("Customer")
	type Request struct {
		Email string `json:"email"`
	}

	// Parse body from JSON to struct
	var user Request
	err := c.BodyParser(&user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var customer models.Customer
	// Check if user credentials are correct
	Customer.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&customer)
	if customer == (models.Customer{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(Message{Msg: "Credenciales incorrectas"})
	}

	// Create token and encrypt it
	encryptedToken, err := pasetoware.CreateToken([]byte(secretSymmetricKey), customer.Name, 12*time.Hour, pasetoware.PurposeLocal)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	type Response struct {
		Token string          `json:"token"`
		User  models.Customer `json:"user"`
	}

	return c.JSON(Response{Token: encryptedToken, User: customer})

}
