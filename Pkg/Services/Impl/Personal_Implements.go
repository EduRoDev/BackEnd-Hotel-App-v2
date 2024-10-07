package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Personal struct{}

func (p Personal) Get() []entities.Personal {
	var personal []entities.Personal
	result := database.Database.Preload("Habitacion").Preload("Habitacion.Habitacion").Find(&personal)
	if result.Error != nil {
		return nil
	}
	return personal
}

func (p Personal) GetID(personal entities.Personal) entities.Personal {
	result := database.Database.Preload("Habitacion").Preload("Habitacion.Habitacion").First(&personal, personal.ID)
	if result.Error != nil {
		return entities.Personal{}
	}
	return personal
}

func (p Personal) Asing(personal entities.Personal) map[string]interface{} {
	result := database.Database.Create(&personal)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear personal")
	}
	return helpers.Success("Personal creada correctamente")
}

func (p Personal) Mod(personal entities.Personal) map[string]interface{} {
	result := database.Database.Model(&personal).Updates(personal)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar personal")
	}
	return helpers.Success("Personal modificada correctamente")
}

func (p Personal) Del(personal entities.Personal) map[string]interface{} {
	result := database.Database.Model(&personal).Delete(&personal)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar personal")
	}
	return helpers.Success("Personal eliminada correctamente")
}
