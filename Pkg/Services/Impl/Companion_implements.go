package impl

import (
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type Companion struct{}

func (c Companion) Get() []entities.Acompañante {
	var acompañante []entities.Acompañante
	result := database.Database.Find(&acompañante)
	if result.Error != nil {
		return nil
	}
	return acompañante
}

func (c Companion) GetID(acompañante entities.Acompañante) entities.Acompañante {
	result := database.Database.First(&acompañante, acompañante.ID)
	if result.Error != nil {
		return entities.Acompañante{}
	}
	return acompañante
}

func (c Companion) Create(acompañante entities.Acompañante) map[string]interface{} {
	result := database.Database.Create(&acompañante)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al crear acompañante")
	}
	return helpers.Success("Acompañante creado correctamente")
}

func (c Companion) Mod(acompañante entities.Acompañante) map[string]interface{} {
	result := database.Database.Model(&acompañante).Updates(acompañante)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al modificar acompañante")
	}
	return helpers.Success("Acompañante modificado correctamente")
}

func (c Companion) Del(acompañante entities.Acompañante) map[string]interface{} {
	result := database.Database.Model(&acompañante).Delete(&acompañante)
	if result.Error != nil {
		return helpers.Error(result.Error, "Error al eliminar acompañante")
	}
	return helpers.Success("Acompañante eliminado correctamente")
}
