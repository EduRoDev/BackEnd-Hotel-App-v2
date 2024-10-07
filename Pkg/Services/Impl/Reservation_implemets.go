package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Reserva struct{}

func (r Reserva) Get() []entities.Reserva {
	var reserva []entities.Reserva
	result := database.Database.Preload("Usuario").Preload("Habitacion").Find(&reserva)
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

func (r Reserva) Create(Reserva entities.Reserva) map[string]interface{} {
	tx := database.Database.Begin()

	var Habitacion entities.Habitacion
	if tx.First(&Habitacion, Reserva.IDHabitacion).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener habitacion")
	}

	if Habitacion.Estado != "disponible" {
		tx.Rollback()
		return helpers.Error(tx.Error, "habitacion no disponible")
	}

	Habitacion.Estado = "reservada"
	if err := tx.Save(&Habitacion).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al actualizar habitacion")
	}

	Reserva.Estado = "pendiente"
	if err := tx.Create(&Reserva).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al crear reserva")
	}

	tx.Commit()
	return helpers.Success("Reserva creada correctamente")
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

func (r Reserva) CancelReserva(reservaId int) map[string]interface{} {
	tx := database.Database.Begin()
	var reserva entities.Reserva
	if tx.First(&reserva, reservaId).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener reserva")
	}
	reserva.Estado = "cancelada"
	if err := tx.Save(&reserva).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al cancelar reserva")
	}

	var Habitacion entities.Habitacion
	if tx.First(&Habitacion, reserva.IDHabitacion).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener habitacion")
	}

	Habitacion.Estado = "disponible"
	if err := tx.Save(&Habitacion).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al modificar habitacion")
	}
	tx.Commit()
	return helpers.Success("Reserva cancelada correctamente")
}
