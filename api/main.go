package main

import (
	"log"

	"github.com/cbsanantero/Routes"
	"github.com/cbsanantero/controllers"
	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	pasetoware "github.com/gofiber/contrib/paseto"
)

const secretSymmetricKey = "symmetric-secret-key (size = 32)"

func main() {
	// Load environment variables from .env file, where API keys and passwords are stored
	if err := godotenv.Load(); err != nil {
		log.Fatalln(".env file err:", err)
	}

	if err := db.ConexionDB(); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer db.DesconectarDB()

	//instancea
	app := fiber.New()
	//cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	
	app.Static("/app", "./public")
	//rutas
	Routes.Rutas(app)

	// Paseto Middleware with local (encrypted) token
	apiGroup := app.Group("api", pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(secretSymmetricKey),
		TokenPrefix:  "",
	}))

	// Restricted Routes
	apiGroup.Get("/restricted", controllers.Restricted)

	//Puerto de escucha
	app.Listen(":3000")
}
