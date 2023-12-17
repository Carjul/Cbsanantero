package main

import (
	"log"
	"os"

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
		log.Println("No .env file found")
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

    // Paseto Middleware with local (encrypted) token
    apiGroup := app.Group("api", pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(secretSymmetricKey),
        TokenPrefix:  "",
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
	//Rutas de Tour
	app.Get("/tour", controllers.GetTour)
	app.Get("/tour/:id", controllers.GetTourById)
	app.Post("/tour", controllers.CreateTour)
	app.Put("/tour/:id", controllers.UpdateTour)
	app.Delete("/tour/:id", controllers.DeleteTour)
	//Rutas de Hoteles
	app.Get("/hotel", controllers.GetHoteles)
	app.Get("/hotel/:id", controllers.GetHotelesById)
	app.Post("/hotel", controllers.CreateHoteles)
	app.Put("/hotel/:id", controllers.UpdateHoteles)
	app.Delete("/hotel/:id", controllers.DeleteHoteles)
	//Rutas de Restaurantes
	app.Get("/restaurante", controllers.GetRestaurante)
	app.Get("/restaurante/:id", controllers.GetRestauranteById)
	app.Post("/restaurante", controllers.CreateRestaurante)
	app.Put("/restaurante/:id", controllers.UpdateRestaurante)
	app.Delete("/restaurante/:id", controllers.DeleteRestaurante)
	//Rutas de Recreacion
	app.Get("/recreacion", controllers.GetRecreacion)
	app.Get("/recreacion/:id", controllers.GetRecreacionById)
	app.Post("/recreacion", controllers.CreateRecreacion)
	app.Put("/recreacion/:id", controllers.UpdateRecreacion)
	app.Delete("/recreacion/:id", controllers.DeleteRecreacion)
	//Rutas de Bar
	app.Get("/bar", controllers.GetBar)
	app.Get("/bar/:id", controllers.GetBarById)
	app.Post("/bar", controllers.CreateBar)
	app.Put("/bar/:id", controllers.UpdateBar)
	app.Delete("/bar/:id", controllers.DeleteBar)
	//Rutas de Artesanias
	app.Get("/artesania", controllers.GetArtesanias)
	app.Get("/artesania/:id", controllers.GetArtesaniasById)
	app.Post("/artesania", controllers.CreateArtesanias)
	app.Put("/artesania/:id", controllers.UpdateArtesanias)
	app.Delete("/artesania/:id", controllers.DeleteArtesanias)
	//Rutas de Comida
	app.Post("/login", controllers.Login)

    // Restricted Routes
    apiGroup.Get("/restricted", controllers.Restricted)
    app.Put("/customerStatus/:id", controllers.UpdateCustomerStatus)
    app.Put("/customerRol/:id", controllers.UpdateCustomerRol)



	//Puerto de escucha
    app.Listen(":"+port)
}