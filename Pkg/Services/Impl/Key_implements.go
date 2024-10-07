package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Llave struct{}

func (l Llave) Get() []entities.Llave {
	var llave []entities.Llave
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").Find(&llave)
	if result.Error != nil {
		return nil
	}
	return llave
}

func (l Llave) GetID(llave entities.Llave) entities.Llave {
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").First(&llave, llave.ID)
	if result.Error != nil {
		return entities.Llave{}
	}
	return llave
}

func (l Llave) Create(llave entities.Llave) map[string]interface{} {
	tx := database.Database.Begin()
	var Reserva entities.Reserva
	if tx.First(&Reserva, llave.IDReserva).Error != nil {
		tx.Rollback()
		return helpers.Error(tx.Error, "Error al obtener reserva")
	}
	if Reserva.Estado != "confirmada" {
		tx.Rollback()
		return helpers.Error(tx.Error, "reserva no confirmada")
	}

	llave.EstadoLlave = "activa"
	if err := tx.Create(&llave).Error; err != nil {
		tx.Rollback()
		return helpers.Error(err, "Error al crear llave")
	}

	tx.Commit()
	return helpers.Success("Llave creada correctamente")
}

func (l Llave) Mod(llave entities.Llave) map[string]interface{} {
	result := database.Database.Model(&llave).Updates(llave)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar llave")
	}
	return helpers.Success("Llave modificada correctamente")
}

func (l Llave) Del(llave entities.Llave) map[string]interface{} {
	result := database.Database.Model(&llave).Delete(&llave)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar llave")
	}
	return helpers.Success("Llave eliminada correctamente")
}


