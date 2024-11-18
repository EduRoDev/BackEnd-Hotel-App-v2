package impl

import (
	"context"
	"errors"
	"log"
	"time"

	auth "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Auth"
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	esquemas "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Esquemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

func (a Admin) Login(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin esquemas.Administrador
	filter := bson.M{"email": email, "password": password}
	err := database.AdminCollection.FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	token, err := auth.GenerateTokenAdmin(admin.Nombre)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a Admin) GetAllAdministradores() ([]esquemas.Administrador, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := database.AdminCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al obtener administradores: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var administradores []esquemas.Administrador

	for cursor.Next(ctx) {
		var admin esquemas.Administrador
		if err := cursor.Decode(&admin); err != nil {
			log.Printf("Error al decodificar administrador: %v", err)
			return nil, err
		}
		administradores = append(administradores, admin)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor: %v", err)
		return nil, err
	}

	return administradores, nil
}

func (a Admin) Create(admin esquemas.Administrador) map[string]interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	admin.ID = primitive.NewObjectID()
	_, err := database.AdminCollection.InsertOne(ctx, admin)
	if err != nil {
		return helpers.Error(err, "Error al crear administrador")
	}

	return helpers.Success("Administrador creado correctamente")
}
