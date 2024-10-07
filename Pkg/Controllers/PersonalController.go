package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	dto "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Dto"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
	impl "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Impl"
	interfaces "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Interfaces"
	"github.com/gorilla/mux"
)

type PersonalController struct {
	l  *log.Logger
	Pl interfaces.Personal
}

func NewPersonalController(l *log.Logger) *PersonalController {
	return &PersonalController{l, &impl.Personal{}}
}

func (pl PersonalController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personal := pl.Pl.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(personal)
}

func (pl PersonalController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	personal := entities.Personal{ID: idStr}
	findPersonal := pl.Pl.GetID(personal)
	if findPersonal.ID == 0 {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonal)
}

func (pl PersonalController) Asign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var personal dto.PersonalDTO
	if err := json.NewDecoder(r.Body).Decode(&personal); err != nil {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Personal{
		ID:       personal.ID,
		Nombre:   personal.Nombre,
		Apellido: personal.Apellido,
		Rol:      personal.Rol,
	}
	findPersonal := pl.Pl.Asing(data)
	if findPersonal["status"] == "error" {
		rp := helpers.ErrorWithStatus("Error in database", "Error al crear personal", "error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findPersonal)
}

func (pl PersonalController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var personal dto.PersonalDTO
	if err := json.NewDecoder(r.Body).Decode(&personal); err != nil {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Personal{
		ID:       idStr,
		Nombre:   personal.Nombre,
		Apellido: personal.Apellido,
		Rol:      personal.Rol,
	}

	findPersonal := pl.Pl.Mod(data)

	if findPersonal["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonal)
}

func (pl PersonalController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Personal{ID: idStr}
	findPersonal := pl.Pl.Del(data)

	if findPersonal["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener personal")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonal)
}
