package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Invoice struct{}

func (i Invoice) Get() []entities.FacturaElectronica {
	var invoice []entities.FacturaElectronica
	result := database.Database.Preload("Pago").Preload("Pago.Reserva").Find(&invoice)
	if result.Error != nil {
		return nil
	}
	return invoice
}

func (i Invoice) GetID(invoice entities.FacturaElectronica) entities.FacturaElectronica {
	result := database.Database.Preload("Pago").Preload("Pago.Reserva").First(&invoice, invoice.ID)
	if result.Error != nil {
		return entities.FacturaElectronica{}
	}
	return invoice
}

func (i Invoice) Asing(invoice entities.FacturaElectronica) map[string]interface{} {
	result := database.Database.Create(&invoice)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear factura")
	}
	return helpers.Success("Factura creada correctamente")
}

func (i Invoice) Mod(invoice entities.FacturaElectronica) map[string]interface{} {
	result := database.Database.Model(&invoice).Updates(invoice)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar factura")
	}
	return helpers.Success("Factura modificada correctamente")
}

func (i Invoice) Del(invoice entities.FacturaElectronica) map[string]interface{} {
	result := database.Database.Model(&invoice).Delete(&invoice)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar factura")
	}
	return helpers.Success("Factura eliminada correctamente")
}
