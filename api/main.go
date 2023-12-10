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
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You must set your 'PORT' environment variable.")
	}

	// Connect to MongoDB
	db.ConexionDB()
	defer db.DesconectarDB()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:"*", 
		AllowMethods:"GET,POST,PUT,DELETE", 
		AllowHeaders:"Origin, Content-Type, Accept",
		AllowCredentials:true, 
	}))

	//Ruta Inicial 
    app.Get("/", controllers.Init)
	//Rutas de Customer
	app.Get("/customer", controllers.GetCustumer)
	app.Get("/customer/:id", controllers.GetCustumerById)
	app.Post("/customer", controllers.CreateCustomer)
	app.Put("/customer/:id", controllers.UpdateCustomer)
	app.Delete("/customer/:id", controllers.DeleteCustomer)
	//Rutas de Trasporte
	app.Get("/trasporte", controllers.GetTrasporte)
	app.Get("/trasporte/:id", controllers.GetTrasporteById)
	app.Post("/trasporte", controllers.CreateTrasporte)
	app.Put("/trasporte/:id", controllers.UpdateTrasporte)
	app.Delete("/trasporte/:id", controllers.DeleteTrasporte)
	//Rutas de Hospedaje
	app.Get("/hospedaje", controllers.GetHospedaje)
	app.Get("/hospedaje/:id", controllers.GetHospedajeById)
	app.Post("/hospedaje", controllers.CreateHospedaje)
	app.Put("/hospedaje/:id", controllers.UpdateHospedaje)
	app.Delete("/hospedaje/:id", controllers.DeleteHospedaje)
	//Puerto de escucha
    app.Listen(":"+port)
}