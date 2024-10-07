package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	Entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type CheckInCheckOut struct{}

func (c CheckInCheckOut) Get() []Entities.CheckInCheckOut {
	var checkInCheckOut []Entities.CheckInCheckOut
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").Find(&checkInCheckOut)
	if result.Error != nil {
		return nil
	}
	return checkInCheckOut
}

func (c CheckInCheckOut) GetID(checkInCheckOut Entities.CheckInCheckOut) Entities.CheckInCheckOut {
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").First(&checkInCheckOut, checkInCheckOut.ID)
	if result.Error != nil {
		return Entities.CheckInCheckOut{}
	}
	return checkInCheckOut
}

func (c CheckInCheckOut) Asing(checkInCheckOut Entities.CheckInCheckOut) map[string]interface{} {
	tx := database.Database.Begin()
	var Reserva Entities.Reserva
	if tx.First(&Reserva, checkInCheckOut.IDReserva).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener reserva")
	}
	if Reserva.Estado != "confirmada" {
		tx.Rollback()
		return helpers.Success("reserva no confirmada")
	}

	checkInCheckOut.Reserva = Reserva
	if err := tx.Create(&checkInCheckOut).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al crear checkInCheckOut")
	}

	tx.Commit()
	return helpers.Success("checkInCheckOut creada correctamente")
}

func (c CheckInCheckOut) Mod(checkInCheckOut Entities.CheckInCheckOut) map[string]interface{} {
	result := database.Database.Model(&checkInCheckOut).Updates(checkInCheckOut)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar checkInCheckOut")
	}
	return helpers.Success("checkInCheckOut modificada correctamente")
}

func (c CheckInCheckOut) Del(checkInCheckOut Entities.CheckInCheckOut) map[string]interface{} {
	result := database.Database.Model(&checkInCheckOut).Delete(&checkInCheckOut)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar checkInCheckOut")
	}
	return helpers.Success("checkInCheckOut eliminada correctamente")
}
