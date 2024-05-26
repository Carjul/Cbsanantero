package controllers

import (
	"context"
	"log"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPedirServicionegocio(c *fiber.Ctx) error {
	service := Instance.Database.Collection("Servicio")

	id := c.Params("id")

	busqueda, err := service.Find(context.TODO(), bson.M{"negocioId": id, "status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el servicio"})
	}

	defer busqueda.Close(context.Background())

	var servicios []bson.M

	if err = busqueda.All(context.Background(), &servicios); err != nil {

		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el servicio"})
	}

	return c.Status(fiber.StatusAccepted).JSON(servicios)

}
func GetPedirServicioClient(c *fiber.Ctx) error {

	service := Instance.Database.Collection("Servicio")

	id := c.Params("id")

	busqueda, err := service.Find(context.TODO(), bson.M{"customer_id":id, "status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el servicio"})
	}

	defer busqueda.Close(context.Background())

	var servicios []bson.M

	if err = busqueda.All(context.Background(), &servicios); err != nil {

		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el servicio"})
	}

	return c.Status(fiber.StatusAccepted).JSON(servicios)

}
func UpdateRevision(c *fiber.Ctx) error {
	type RevisionStatus struct {
		Revision bool `json:"revision,omitempty" bson:"revision,omitempty"`
	}
	service := Instance.Database.Collection("Servicio")

	id := c.Params("id")
	ObjID, errid := primitive.ObjectIDFromHex(id)
	if errid != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Id de servicio no valido"})
	}
	var revisionStatus RevisionStatus

	if err := c.BodyParser(&revisionStatus); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Message{Msg: "Error al leer la data"})
	}

	_, err := service.UpdateOne(context.Background(), bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"revision": revisionStatus.Revision}})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el bar"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Se actualizo la revision"})

}

func CreatePedirServicio(c *fiber.Ctx) error {
	service := Instance.Database.Collection("Servicio")
	customer := Instance.Database.Collection("Customer")
	artesanias := Instance.Database.Collection("Artesanias")
	db := Instance.Database
	type n struct {
		Name    string `json:"name,omitempty" bson:"name,omitempty"`
		Address string `json:"address,omitempty" bson:"address,omitempty"`
		Image   string `json:"image,omitempty" bson:"image,omitempty"`
	}

	var servicio models.ClientePeticion
	var Negocio n
	//query params
	TipoNegocio := c.Query("tipo")

	if err := c.BodyParser(&servicio); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Message{Msg: "Error al leer la data"})
	}

	servicio.Status = "Activo"
	servicio.Revision = false
	servicio.Fecha = config.GetDate()
	servicio.Hora = config.GetHour()
	
	_, err := service.InsertOne(context.Background(), servicio)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el servicio"})
	}

	idCustomer, err := primitive.ObjectIDFromHex(servicio.CustomerID)																																																											
	if err != nil {
		log.Fatal(err)
	}
	var customers models.Customer
	err = customer.FindOne(context.TODO(), bson.M{"_id": idCustomer}).Decode(&customers)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})
	}

	objIDNegocio, err := primitive.ObjectIDFromHex(servicio.NegocioId)


	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el servicio"})
	}
	switch TipoNegocio {
	case "Artesania":
		err = artesanias.FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	case "Bar":
		err = db.Collection("Bares").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	case "Hospedaje":
		err = db.Collection("Hospedaje").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})

		}
	case "Hotel":
		err = db.Collection("Hoteles").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	case "Recreacion":
		err = db.Collection("Recreacion").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	case "Restaurante":
		err = db.Collection("Restaurantes").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	case "Tour":
		err = db.Collection("Tour").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}

	case "Transporte":
		err = db.Collection("Traporte").FindOne(context.Background(), bson.M{"_id": objIDNegocio, "status": "Activo"}).Decode(&Negocio)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el negocio"})
		}
	}

	//send email
	config.GoMail(customers.Email, "Petición de servicio a "+Negocio.Name, `
	<html>
	<head>
		<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
		}
		h3 {
			color: blue;
		}
		img {
			width: 100%;
			height: auto;
		}
		p {
			color: black;
		}
		</style>
	</head>
	<body>
		<h3>Solucitud de servicio para `+Negocio.Name+`</h3>
		 
		<p> Solicitante: `+servicio.Nombre+`</p>
		<p> Telefono: `+servicio.Celular+`</p>
		<p> Correo: `+servicio.Correo+`</p>
		<p> Descripcion: `+servicio.Description+`</p>
		<p> Fecha: `+servicio.Fecha+`</p>
		<p> Hora: `+servicio.Hora+`</p>


	</body>
	</html>
`)
	config.GoMail(servicio.Correo, "Solucitud de servicio enviada a "+Negocio.Name, `
	<html>
	<head>
		<style>
			
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
			}
			h3 {
				color: blue;
			}
			img {
				width: 40%;
				height: auto;
			}
			p {
				color: black;
			}
		</style>
	</head>
	<body>
	
		<h3>`+Negocio.Name+`</h3>
		 
		<p> Direcion: `+Negocio.Address+`</p>
		
		<img src=`+Negocio.Image+` alt="Artesania Image">

		<p> Se ha enviado una solicitud de servicio al <strong>administrador</strong> del negocio
		`+Negocio.Name+`, favor esperar su respuesta.
		</p>

	</body>
	</html>
`)

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Mensaje enviado correctamente"})

}
func DeleteServicio(c *fiber.Ctx) error {

	service := Instance.Database.Collection("Servicio")

	id := c.Params("id")

	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "Id de servicio no valido"})
	}

	_, err = service.DeleteOne(context.Background(), bson.M{"_id": ObjID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el servicio"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Se elimino el servicio"})

}
