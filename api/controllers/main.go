package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Msg string `json:"message"`
}

func Init(c *fiber.Ctx) error {
	
	return c.Status(fiber.StatusOK).JSON(Message{
		Msg: "Bienvenido a la API de Go Fiber",
	})
}