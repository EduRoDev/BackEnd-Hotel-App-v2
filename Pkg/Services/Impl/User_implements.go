package impl

import (
	"fmt"

	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
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

func (u User) Asing(User entities.Usuario) map[string]interface{} {
	result := database.Database.Create(&User)
	if result.Error != nil {
		return map[string]interface{}{"error": result.Error, "message": "Error al crear usuario"}
	}
	return map[string]interface{}{"success": "Validacion correcta", "message": "Usuario creado correctamente"}
}

func (u User) Mod(User entities.Usuario) map[string]interface{} {
	result := database.Database.Model(&User).Updates(User)
	if result.Error != nil {
		return map[string]interface{}{"error": result.Error, "message": "Error al modificar usuario"}
	}
	return map[string]interface{}{"success": "Validacion correcta", "message": "Usuario modificado correctamente"}
}

func (u User) Del(user entities.Usuario) map[string]interface{} {
	tx := database.Database.Begin()
	result := tx.Delete(&user)
	if result.Error != nil {
		tx.Rollback()
		return map[string]interface{}{"error": result.Error, "message": "Error al eliminar usuario"}
	}

	var usuarios []entities.Usuario
	if err := tx.Order("id asc").Find(&usuarios).Error; err != nil {
		tx.Rollback()
		return map[string]interface{}{"error": err, "message": "Error al obtener usuarios"}
	}

	for i, u := range usuarios {
		newID := uint(i + 1)
		if err := tx.Model(&u).Update("id", newID).Error; err != nil {
			tx.Rollback()
			return map[string]interface{}{"error": err, "message": "Error al reasignar IDs"}
		}
	}

	nextID := len(usuarios) + 1
	query := fmt.Sprintf("ALTER TABLE usuarios AUTO_INCREMENT = %d", nextID)
	if err := tx.Exec(query).Error; err != nil {
		tx.Rollback()
		return map[string]interface{}{"error": err, "message": "Error al reiniciar el autoincremental"}
	}

	tx.Commit()
	return map[string]interface{}{"success": true, "message": "Usuario eliminado y IDs reasignados correctamente"}
}
