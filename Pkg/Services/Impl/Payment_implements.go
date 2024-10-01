package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Pago struct{}

func (p Pago) Get() []entities.Pago {
	var pago []entities.Pago
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").Find(&pago)
	if result.Error != nil {
		return nil
	}
	return pago
}

func (p Pago) GetID(pago entities.Pago) entities.Pago {
	result := database.Database.Preload("Reserva").Preload("Reserva.Usuario").Preload("Reserva.Habitacion").First(&pago, pago.ID)
	if result.Error != nil {
		return entities.Pago{}
	}
	return pago
}

func (p Pago) Create(Pago entities.Pago) map[string]interface{} {
	result := database.Database.Create(&Pago)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al asignar pago")
	}
	return helpers.Success("Pago asignado correctamente")
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
