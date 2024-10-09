package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type PersonalRoom struct{}

func (p PersonalRoom) Get() []entities.PersonalHabitacion {
	var personalRoom []entities.PersonalHabitacion
	result := database.Database.Preload("Habitacion").Preload("Personal").Find(&personalRoom)
	if result.Error != nil {
		return nil
	}
	return personalRoom
}

func (p PersonalRoom) GetID(personalRoom entities.PersonalHabitacion) entities.PersonalHabitacion {
	result := database.Database.Preload("Habitacion").Preload("Personal").First(&personalRoom, personalRoom.ID)
	if result.Error != nil {
		return entities.PersonalHabitacion{}
	}
	return personalRoom
}

func (p PersonalRoom) Asing(personalRoom entities.PersonalHabitacion) map[string]interface{} {
	tx := database.Database.Begin()
	var Habitacion entities.Habitacion
	if tx.First(&Habitacion, personalRoom.IDHabitacion).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener habitacion")
	}
	if Habitacion.Estado != "disponible" {
		tx.Rollback()
		return helpers.Error(tx.Error, "habitacion no disponible")
	}

	personalRoom.Habitacion = Habitacion
	if err := tx.Create(&personalRoom).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al asignar personal a la habitacion")
	}

	tx.Commit()
	return helpers.Success("Personal asignado correctamente")
}

func (p PersonalRoom) Mod(personalRoom entities.PersonalHabitacion) map[string]interface{} {
	result := database.Database.Model(&personalRoom).Updates(personalRoom)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar personalRoom")
	}
	return helpers.Success("PersonalRoom modificada correctamente")
}

func (p PersonalRoom) Del(personalRoom entities.PersonalHabitacion) map[string]interface{} {
	result := database.Database.Model(&personalRoom).Delete(&personalRoom)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar personalRoom")
	}
	return helpers.Success("PersonalRoom eliminada correctamente")
}
