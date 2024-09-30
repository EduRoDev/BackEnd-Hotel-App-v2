package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Reserva struct{}

func (r Reserva) Get() []entities.Reserva {
	var reserva []entities.Reserva
	result := database.Database.Find(&reserva).Preload("usuarios", "habitacions")
	if result.Error != nil {
		return nil
	}
	return reserva
}

func (r Reserva) GetID(reserva entities.Reserva) entities.Reserva {
	result := database.Database.Preload("Usuario").Preload("Habitacion").First(&reserva, reserva.ID)
	if result.Error != nil {
		return entities.Reserva{}
	}
	return reserva
}


func (r Reserva) Asing(Reserva entities.Reserva) map[string]interface{} {
	result := database.Database.Create(&Reserva)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al asignar reserva")
	}
	return helpers.Success("Reserva asignada correctamente")
}

func (r Reserva) Mod(Reserva entities.Reserva) map[string]interface{} {
	result := database.Database.Model(&Reserva).Updates(Reserva)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar reserva")
	}
	return helpers.Success("Reserva modificada correctamente")
}

func (r Reserva) Del(reserva entities.Reserva) map[string]interface{} {
	result := database.Database.Model(&reserva).Delete(&reserva)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar reserva")
	}
	return helpers.Success("Reserva eliminada correctamente")
}
