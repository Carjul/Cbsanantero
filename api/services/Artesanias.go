package services

import (
	"context"
	"time"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Msg string
}
type Artesanias struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Image   string             `json:"image,omitempty" bson:"image,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Customer struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Image    string             `json:"image,omitempty" bson:"image,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Rol      string             `json:"rol,omitempty" bson:"rol,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
}

func GetArtesanias() interface{}  {
	artesanias := db.Artesanias
	busqueda, err := artesanias.Find(context.TODO(), bson.M{"status":"Activo"})
	if err != nil {
		return  "error al buscar artesania"
	}
	defer busqueda.Close(context.Background())

	var artesania []bson.M
	if err = busqueda.All(context.Background(), &artesania); err != nil {
		return  "No se pudo encontrar la artesania"
	}
	return artesania
}

func GetArtesaniasById(id string) interface{} {
	artesanias := db.Artesanias
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No existe el ID de la artesania"}
	}
	busqueda := artesanias.FindOne(context.Background(), bson.M{"_id": objID,"status":"Activo"})
	var artesania Artesanias
	if err := busqueda.Decode(&artesania); err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}
	return artesania
}

func CreateArtesanias(data *Artesanias) interface{} {
	artesanias := db.Artesanias
	customer := db.Customer
	idc:= data.CustomerID
	objID, err := primitive.ObjectIDFromHex(idc)

	go config.UploadImageLocal(data.Image)
	time.Sleep(1 * time.Second)
	url := config.UploadImage()
	if url == "error al subir la imagen a cloudinary"{
		return Message{Msg: "error al subir la imagen a cloudinary"}
	}
	if err != nil {
		return Message{Msg: "El _id Customer no es valido"}
	}
	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})
	var customerData Customer
	if err = busqueda.Decode(&customerData); err != nil {
		return Message{Msg: "No se pudo encontrar el cliente"}
	}
	_, err = artesanias.InsertOne(context.Background(), data)
	if err != nil {
		return Message{Msg: "No se pudo crear la artesania"}
	}
	return Message{Msg: "Artesania creada"}
}

func UpdateArtesanias(data *Artesanias, id string) interface{} {
	artesanias := db.Artesanias
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}
	if data.Image != "" {
		go config.UploadImageLocal(data.Image)
		time.Sleep(1 * time.Second)
		url := config.UploadImage()
		if url == "error al subir la imagen a cloudinary"{
			return Message{Msg: "error al subir la imagen a cloudinary"}
		}
		data.Image = url
	
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
	artesanias := db.Artesanias
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Message{Msg: "No se pudo encontrar la artesania"}
	}
	result, err := artesanias.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return Message{Msg: "error al eliminar la artesania"}
	}
	if result.DeletedCount == 0 {
		return Message{Msg: "No se pudo eliminar la artesania"}
	}
	return Message{Msg: "Artesania eliminada"}
}