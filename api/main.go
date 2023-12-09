package main

import (
	"log"
	"os"

	"github.com/cbsanantero/controllers"
	"github.com/cbsanantero/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file, where API keys and passwords are stored
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// Connect to MongoDB
	db.ConexionDB()
	defer db.DesconectarDB()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", 
		AllowMethods:     "GET,POST,PUT,DELETE", 
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true, 
	}))

	//Ruta Inicial 
    app.Get("/", controllers.Init)
	
	app.Get("/customer", controllers.GetCustumer)
	app.Get("/customer/:id", controllers.GetCustumerById)
	app.Post("/customer", controllers.CreateCustomer)

    app.Listen(":3000")
}