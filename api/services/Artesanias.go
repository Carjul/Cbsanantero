package services

import (
	"context"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Msg string
}

func GetArtesanias() interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	busqueda, err := artesanias.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return "error al buscar artesania"
	}
	defer busqueda.Close(context.Background())

	var artesania []bson.M
	if err = busqueda.All(context.Background(), &artesania); err != nil {
		return "No se pudo encontrar la artesania"
	}
	return artesania
}

func GetArtesaniasById(id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No existe el ID de la artesania"}
	}
	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID, "status": "Activo"})
	var artesania models.Artesanias
	if err := busqueda.Decode(&artesania); err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}
	return artesania
}

func CreateArtesanias(data *models.Artesanias) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	customer := Instance.Database.Collection("Customer")
	idc := data.CustomerID
	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return Message{Msg: "El _id Customer no es valido"}
	}
	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})
	var customerData models.Customer
	if err = busqueda.Decode(&customerData); err != nil {
		return Message{Msg: "No se pudo encontrar el cliente"}
	}
	_, err = artesanias.InsertOne(context.Background(), data)
	if err != nil {
		return Message{Msg: "No se pudo crear la artesania"}
	}
	return Message{Msg: "Artesania creada"}

}

func UpdateArtesanias(data *models.Artesanias, id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}

	result, err := artesanias.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return Message{Msg: "error al actualizar la artesania"}
	}
	if result.ModifiedCount == 0 {
		return Message{Msg: "No se pudo actualizar la artesania"}
	}
	return Message{Msg: "Artesania actualizada"}
}

func DeleteArtesanias(id string) interface{} {
	artesanias := Instance.Database.Collection("Artesanias")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "_id de la artesania no valido"}
	}
	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID})
	var artesania models.Artesanias
	if err := busqueda.Decode(&artesania); err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}

	result, err := artesanias.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return Message{Msg: "error al eliminar la artesania"}
	}
	if result.DeletedCount == 0 {
		return Message{Msg: "No se pudo eliminar la artesania"}
	}
	config.DeleteImage(artesania.Image)
	return Message{Msg: "Artesania eliminada"}
}
