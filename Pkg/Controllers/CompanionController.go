package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
	impl "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Impl"
	interfaces "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Interfaces"
	"github.com/gorilla/mux"
)

type CompanionController struct {
	l *log.Logger
	C interfaces.Acompañante
}

func NewCompanionController(l *log.Logger) *CompanionController {
	return &CompanionController{l, &impl.Companion{}}
}

func (c CompanionController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	acompañante := c.C.Get()
	if acompañante == nil {
		w.WriteHeader(http.StatusNotFound)
		helpers.Error(nil, "No se encontraron acompañantes")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(acompañante)
}

func (c CompanionController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		helpers.Error(err, "No se encontraron acompañantes")
		return
	}
	acompañante := entities.Acompañante{ID: idStr}
	findAcompañante := c.C.GetID(acompañante)
	if findAcompañante.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		helpers.Error(nil, "No se encontraron acompañantes")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findAcompañante)
}

func (c CompanionController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var acompañante entities.Acompañante
	if err := json.NewDecoder(r.Body).Decode(&acompañante); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Error(err, "Error al obtener acompañante")
		return
	}
	findAcompañante := c.C.Create(acompañante)
	if findAcompañante["error"] != nil {
		w.WriteHeader(http.StatusInternalServerError)
		helpers.Error(nil, "Error al crear acompañante")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findAcompañante)
}

func (c CompanionController) Mod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		helpers.Error(err, "No se encontraron acompañantes")
		return
	}
	acompañante := entities.Acompañante{ID: idStr}
	if err := json.NewDecoder(r.Body).Decode(&acompañante); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Error(err, "Error al obtener acompañante")
		return
	}
	findAcompañante := c.C.Mod(acompañante)
	if findAcompañante["error"] != nil {
		w.WriteHeader(http.StatusInternalServerError)
		helpers.Error(nil, "Error al modificar acompañante")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findAcompañante)
}

func (c CompanionController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		helpers.Error(err, "No se encontraron acompañantes")
		return
	}
	acompañante := entities.Acompañante{ID: idStr}
	findAcompañante := c.C.Del(acompañante)
	if findAcompañante["error"] != nil {
		w.WriteHeader(http.StatusInternalServerError)
		helpers.Error(nil, "Error al eliminar acompañante")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findAcompañante)
}