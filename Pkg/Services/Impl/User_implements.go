package impl

import (
	"errors"

	auth "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Auth"
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type User struct{}

func (u User) Login(nombre string, numeroDocumento string) (string, error) {
	var usuario entities.Usuario

	result := database.Database.Where("nombre = ? AND numero_documento = ?", nombre, numeroDocumento).First(&usuario)
	if result.Error != nil {
		return "", errors.New("usuario no encontrado")
	}

	token, err := auth.GenerateToken(usuario.ID, usuario.Nombre)
	if err != nil {
		return "", errors.New("error al generar el token")
	}
	return token, nil
}

func (u User) Get() []entities.Usuario {
	var user []entities.Usuario
	result := database.Database.Find(&user)
	if result.Error != nil {
		return nil
	}
	return user
}

func (u User) GetID(User entities.Usuario) entities.Usuario {
	result := database.Database.First(&User, User.ID)
	if result.Error != nil {
		return entities.Usuario{}
	}
	return User
}

func (u User) GetUser(nombre string) entities.Usuario {
	var User entities.Usuario
	result := database.Database.Where("nombre = ?", nombre).First(&User)
	if result.Error != nil {
		return entities.Usuario{}
	}
	return User
}

func (u User) Create(User entities.Usuario) map[string]interface{} {
	result := database.Database.Create(&User)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear usuario")
	}
	return helpers.Success("Usuario creado correctamente")
}

func (u User) Mod(User entities.Usuario) map[string]interface{} {
	result := database.Database.Model(&User).Updates(User)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar usuario")
	}
	return helpers.Success("Usuario modificado correctamente")
}

func (u User) Del(user entities.Usuario) map[string]interface{} {
	result := database.Database.Model(&user).Delete(&user)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar usuario")
	}
	return helpers.Success("Usuario eliminado correctamente")
}
