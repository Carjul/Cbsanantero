package Routes

import (
	"github.com/cbsanantero/controllers"
	"github.com/gofiber/fiber/v2"
)

func Rutas(app *fiber.App) {
	//Ruta Inicial
	app.Get("/app", controllers.Init)
	//Rutas de Customer
	app.Get("/app/customer", controllers.GetCustumer)
	app.Get("/app/customer/:id", controllers.GetCustumerById)
	app.Post("/app/customer", controllers.CreateCustomer)
	app.Put("/app/customer/:id", controllers.UpdateCustomer)
	app.Put("/app/customerStatus/:id", controllers.UpdateCustomerStatus)
	app.Put("/app/customerRol/:id", controllers.UpdateCustomerRol)
	app.Delete("/app/customer/:id", controllers.DeleteCustomer)
	//Rutas de Trasporte
	app.Get("/app/trasporte", controllers.GetTrasporte)
	app.Get("/app/trasporte/:id", controllers.GetTrasporteById)
	app.Post("/app/trasporte", controllers.CreateTrasporte)
	app.Put("/app/trasporte/:id", controllers.UpdateTrasporte)
	app.Delete("/app/trasporte/:id", controllers.DeleteTrasporte)
	//Rutas de Hospedaje
	app.Get("/app/hospedaje", controllers.GetHospedaje)
	app.Get("/app/hospedaje/:id", controllers.GetHospedajeById)
	app.Post("/app/hospedaje", controllers.CreateHospedaje)
	app.Put("/app/hospedaje/:id", controllers.UpdateHospedaje)
	app.Delete("/app/hospedaje/:id", controllers.DeleteHospedaje)
	//Rutas de Tour
	app.Get("/app/tour", controllers.GetTour)
	app.Get("/app/tour/:id", controllers.GetTourById)
	app.Post("/app/tour", controllers.CreateTour)
	app.Put("/app/tour/:id", controllers.UpdateTour)
	app.Delete("/app/tour/:id", controllers.DeleteTour)
	//Rutas de Hoteles
	app.Get("/app/hotel", controllers.GetHoteles)
	app.Get("/app/hotel/:id", controllers.GetHotelesById)
	app.Post("/app/hotel", controllers.CreateHoteles)
	app.Put("/app/hotel/:id", controllers.UpdateHoteles)
	app.Delete("/app/hotel/:id", controllers.DeleteHoteles)
	//Rutas de Restaurantes
	app.Get("/app/restaurante", controllers.GetRestaurante)
	app.Get("/app/restaurante/:id", controllers.GetRestauranteById)
	app.Post("/app/restaurante", controllers.CreateRestaurante)
	app.Put("/app/restaurante/:id", controllers.UpdateRestaurante)
	app.Delete("/app/restaurante/:id", controllers.DeleteRestaurante)
	//Rutas de Recreacion
	app.Get("/app/recreacion", controllers.GetRecreacion)
	app.Get("/app/recreacion/:id", controllers.GetRecreacionById)
	app.Post("/app/recreacion", controllers.CreateRecreacion)
	app.Put("/app/recreacion/:id", controllers.UpdateRecreacion)
	app.Delete("/app/recreacion/:id", controllers.DeleteRecreacion)
	//Rutas de Bar
	app.Get("/app/bar", controllers.GetBar)
	app.Get("/app/bar/:id", controllers.GetBarById)
	app.Post("/app/bar", controllers.CreateBar)
	app.Put("/app/bar/:id", controllers.UpdateBar)
	app.Delete("/app/bar/:id", controllers.DeleteBar)
	//Rutas de Artesanias
	app.Get("/app/artesania", controllers.GetArtesanias)
	app.Get("/app/artesania/:id", controllers.GetArtesaniasById)
	app.Post("/app/artesania", controllers.CreateArtesanias)
	app.Put("/app/artesania/:id", controllers.UpdateArtesanias)
	app.Delete("/app/artesania/:id", controllers.DeleteArtesanias)
	//galeria
	app.Get("/app/galeria/:id", controllers.GetGaleria)
	app.Post("/app/galeria", controllers.CreateGaleria)
	app.Put("/app/galeria/:id", controllers.UpdateGaleria)
	//pedir servicio
	app.Post("/app/pedirServicio", controllers.CreatePedirServicio)
	app.Get("/app/pedirServicion/:id", controllers.GetPedirServicionegocio)
	app.Get("/app/pedirServicioc/:id", controllers.GetPedirServicioClient)
	app.Put("/app/ServicioStatus/:id", controllers.UpdateRevision)
	app.Delete("/app/pedirServicio/:id", controllers.DeleteServicio)

	//Rutas login
	app.Post("/app/login", controllers.Login)
	app.Post("/app/loginGoogle", controllers.LoginG)

}
