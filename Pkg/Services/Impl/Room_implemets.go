package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Habitacion struct{}

func (h Habitacion) Get() []entities.Habitacion {
	var habitacion []entities.Habitacion
	result := database.Database.Find(&habitacion)
	if result.Error != nil {
		return nil
	}
	return habitacion
}

func (h Habitacion) GetID(Habitacion entities.Habitacion) entities.Habitacion {
	result := database.Database.First(&Habitacion, Habitacion.ID)
	if result.Error != nil {
		return entities.Habitacion{}
	}
	return Habitacion
}

func (h Habitacion) Create(Habitacion entities.Habitacion) map[string]interface{} {
	result := database.Database.Create(&Habitacion)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear habitacion")
	}
	return helpers.Success("Habitacion creado correctamente")
}

func (h Habitacion) Mod(Habitacion entities.Habitacion) map[string]interface{} {
	result := database.Database.Model(&Habitacion).Updates(Habitacion)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar habitacion")
	}
	return helpers.Success("Habitacion modificado correctamente")
}

func (h Habitacion) Del(habitacion entities.Habitacion) map[string]interface{} {
	result := database.Database.Model(&habitacion).Delete(&habitacion)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar habitacion")
	}
	return helpers.Success("Habitacion eliminado correctamente")
}
