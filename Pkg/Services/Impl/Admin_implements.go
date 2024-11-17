package impl

import (
	auth "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Auth"
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Admin struct{}

func (a Admin) Login(email string, password string) (string, error) {
	var admin entities.Administrador
	result := database.Database.First(&admin, "email = ? and password = ?", email, password)
	if result.Error != nil {
		return "", result.Error
	}
	if admin.Password != password {
		return "", result.Error
	}
	token, err := auth.GenerateTokenAdmin(admin.Nombre)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a Admin) Create(admin entities.Administrador) map[string]interface{} {
	result := database.Database.Create(&admin)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear administrador")
	}
	return helpers.Success("Administrador creado correctamente")
}
