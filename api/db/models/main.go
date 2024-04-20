package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientePeticion struct {
	Nombre      string `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Celular     string `json:"celular,omitempty" bson:"celular,omitempty"`
	Correo      string `json:"correo,omitempty" bson:"correo,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Revision    bool
	Fecha       string `json:"fecha,omitempty" bson:"fecha,omitempty"`
	Hora        string `json:"hora,omitempty" bson:"hora,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty"`
	ClienteId   string `json:"clienteId,omitempty" bson:"clienteId,omitempty"`
	CustomerId  string `json:"customerId,omitempty" bson:"customerId,omitempty"`
	NegocioId   string `json:"negocioId,omitempty" bson:"negocioId,omitempty"`
}

type Artesanias struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Galeria struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Photos     []string           `json:"photos,omitempty" bson:"photos,omitempty"`
	NegocioID  string             `json:"negocio_id,omitempty" bson:"negocio_id,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Bar struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Customer struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Rol         string             `json:"rol,omitempty" bson:"rol,omitempty"`
	TipoNegocio string             `json:"tipo_negocio,omitempty" bson:"tipo_negocio,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
}

type Hospedaje struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Price       string             `json:"price,omitempty" bson:"price,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Hoteles struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	Image      string             `json:"image,omitempty" bson:"image,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Price      string             `json:"price,omitempty" bson:"price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Recreacion struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	Services   string             `json:"services,omitempty" bson:"services,omitempty"`
	Image      string             `json:"image,omitempty" bson:"image,omitempty"`
	Price      string             `json:"price,omitempty" bson:"price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Restaurante struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID  string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Tour struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Image      string             `json:"image,omitempty" bson:"image,omitempty"`
	Trasporte  string             `json:"trasporte,omitempty" bson:"trasporte,omitempty"`
	Price      string             `json:"price,omitempty" bson:"price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type Trasporte struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tipo       string             `json:"tipo,omitempty" bson:"tipo,omitempty"`
	Image      string             `json:"image,omitempty" bson:"image,omitempty"`
	Placa      string             `json:"placa,omitempty" bson:"placa,omitempty"`
	Conductor  string             `json:"conductor,omitempty" bson:"conductor,omitempty"`
	Celular    string             `json:"celular,omitempty" bson:"celular,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerID string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}
