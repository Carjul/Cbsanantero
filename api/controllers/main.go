package controllers

import (
	"github.com/cbsanantero/config"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Msg string `json:"message"`
}

func Init(c *fiber.Ctx) error {
	config.GoMail()
	return c.Status(fiber.StatusOK).JSON(Message{
		Msg: "Bienvenido a la API de Costa Brisa",
	})
}
