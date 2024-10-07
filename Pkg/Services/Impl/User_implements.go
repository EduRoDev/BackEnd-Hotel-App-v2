package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type User struct{}

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

func (u User) LastID(User entities.Usuario) entities.Usuario{
	result := database.Database.Last(&User)
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
