package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var AdminCollection *mongo.Collection

func Connect() {
	uri := "mongodb://localhost:27017/"
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	log.Println("Conexión exitosa a MongoDB local")
	Client = client

	AdminCollection = client.Database("administradores").Collection("administradores")
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error cerrando conexión a MongoDB: %v", err)
	}
	log.Println("Conexión cerrada a MongoDB")
}
