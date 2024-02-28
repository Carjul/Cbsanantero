package main

import (
	"log"
	"os"

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
		log.Println("No .env file found")
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You must set your 'PORT' environment variable.")
	}

	// Connect to MongoDB
	db.ConexionDB()
	defer db.DesconectarDB()

	//instancea
    app := fiber.New()

	//rutas
	Routes.Rutas(app)

	//cors
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
	apiGroup.Use(
		cors.New(cors.Config{
			AllowOrigins:"*", 
			AllowMethods:"GET,POST,PUT,DELETE", 
			AllowHeaders:"Origin, Content-Type, Accept",
			AllowCredentials:true, 
		}), // Add a comma here
	)
    // Restricted Routes
    apiGroup.Get("/restricted", controllers.Restricted)
	

	//Puerto de escucha
    app.Listen(":"+port)
}