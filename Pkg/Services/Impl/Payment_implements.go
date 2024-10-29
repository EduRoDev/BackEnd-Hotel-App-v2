package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Pago struct{}

func (p Pago) Get() []entities.Pago {
	var pago []entities.Pago
	result := database.Database.Preload("Reserva").
		Preload("Reserva.Usuario").
		Preload("Reserva.Usuario.Acompañante").
		Preload("Reserva.Habitaciones.Habitacion").
		Find(&pago)
	if result.Error != nil {
		return nil
	}
	return pago
}

func (p Pago) GetID(pago entities.Pago) entities.Pago {
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitaciones.Habitacion").First(&pago, pago.ID)
	if result.Error != nil {
		return entities.Pago{}
	}
	return pago
}

func (p Pago) Create(Pago entities.Pago) map[string]interface{} {
	tx := database.Database.Begin()

	
	var Reserva entities.Reserva
	if err := tx.First(&Reserva, Pago.IDReserva).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al obtener reserva")
	}

	if Reserva.Estado != "pendiente" {
		tx.Rollback()
		return helpers.Error(tx.Error, "reserva no pendiente")
	}

	Reserva.Estado = "confirmada"
	if err := tx.Save(&Reserva).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al actualizar reserva")
	}

	var habitacion entities.Habitacion
	if err := tx.First(&habitacion, Reserva.IDHabitacion).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al obtener habitacion")
	}

	habitacion.Estado = "ocupada"
	if err := tx.Save(&habitacion).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al actualizar habitacion")
	}

	if err := tx.Create(&Pago).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al crear el pago")
	}
	
	tx.Commit()
	return helpers.Success("Pago creado correctamente")
}

func (p Pago) Mod(Pago entities.Pago) map[string]interface{} {
	result := database.Database.Model(&Pago).Updates(Pago)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar pago")
	}
	return helpers.Success("Pago modificado correctamente")
}

func (p Pago) Del(pago entities.Pago) map[string]interface{} {
	result := database.Database.Model(&pago).Delete(&pago)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar pago")
	}
	return helpers.Success("Pago eliminado correctamente")
}

func (p Pago) Cancel(idReserva int) map[string]interface{} {
	tx := database.Database.Begin()
	var reserva entities.Reserva
	if tx.First(&reserva, idReserva).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener reserva")
	}
	reserva.Estado = "cancelada"
	if err := tx.Save(&reserva).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al cancelar reserva")
	}

	var Pago entities.Pago
	if tx.First(&Pago, idReserva).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener pago")
	}

	Pago.Estado = "cancelada"
	if err := tx.Save(&Pago).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al modificar pago")
	}
	tx.Commit()
	return helpers.Success("Pago cancelado correctamente")
}