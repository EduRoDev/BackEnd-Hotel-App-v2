package impl

import (
	"time"

	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Reserva struct{}

func (r Reserva) Get() []entities.Reserva {
	var reservas []entities.Reserva
	result := database.Database.
		Preload("Usuario").
		Preload("Usuario.Acompañante").
		Preload("Habitacion").
		Find(&reservas)
	if result.Error != nil {
		return nil
	}
	return reservas
}

func (r Reserva) GetID(reserva entities.Reserva) entities.Reserva {
	result := database.Database.
		Preload("Usuario").
		Preload("Usuario.Acompañante").
		Preload("Habitacion").
		First(&reserva, reserva.ID)
	if result.Error != nil {
		return entities.Reserva{}
	}
	return reserva
}

func (r Reserva) GetByUsuarioYFecha(idUsuario int, fechaEntrada time.Time) entities.Reserva {
	var reserva entities.Reserva
	result := database.Database.
		Preload("Usuario").
		Preload("Usuario.Acompañante").
		Preload("Habitacion").
		Where("id_usuario = ? AND fecha_entrada = ?", idUsuario, fechaEntrada).
		Find(&reserva)
	if result.Error != nil {
		return entities.Reserva{}
	}
	return reserva
}

func (r Reserva) Create(Reserva entities.Reserva) map[string]interface{} {
	tx := database.Database.Begin()

	var Habitacion entities.Habitacion
	if tx.Set("gorm:query_option", "FOR UPDATE").First(&Habitacion, Reserva.IDHabitacion).Error != nil {
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

