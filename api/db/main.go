package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Customer *mongo.Collection
var Traporte *mongo.Collection
var Hospedaje *mongo.Collection
var Tour *mongo.Collection
var Hoteles *mongo.Collection
var Restaurantes *mongo.Collection
var Recreacion *mongo.Collection
var Bares *mongo.Collection
var Artesanias *mongo.Collection
var Galeria *mongo.Collection

func ConexionDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return fmt.Errorf("debes configurar la variable de entorno 'MONGODB_URI'")
	}

	var err error
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error al conectar con MongoDB: %v", err)
	} else {
		log.Println("Conexión exitosa a MongoDB")
	}

	// Verificar la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = Client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("no se pudo realizar el ping a la base de datos: %v", err)
	}

	Db := Client.Database("Cbsanantero")

	Customer = Db.Collection("Customer")
	Traporte = Db.Collection("Trasporte")
	Hospedaje = Db.Collection("Hospedaje")
	Tour = Db.Collection("Tour")
	Hoteles = Db.Collection("Hoteles")
	Restaurantes = Db.Collection("Restaurantes")
	Recreacion = Db.Collection("Recreacion")
	Bares = Db.Collection("Bares")
	Artesanias = Db.Collection("Artesanias")
	Galeria = Db.Collection("Galeria")

	return nil
}

func DesconectarDB() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Printf("Error al desconectar la base de datos: %v", err)
		} else {
			log.Println("Desconexión exitosa de MongoDB")
		}
	}
}
