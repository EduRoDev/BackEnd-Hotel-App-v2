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

type RoomController struct {
	l  *log.Logger
	Hb interfaces.Habitacion
}

func NewRoomController(l *log.Logger) *RoomController {
	return &RoomController{l, &impl.Habitacion{}}
}

func (rm RoomController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	habitacion := rm.Hb.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(habitacion)
}

func (rm RoomController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	habitacion := entities.Habitacion{ID: idStr}
	findHabitacion := rm.Hb.GetID(habitacion)

	if findHabitacion.ID == 0 {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findHabitacion)
}

func (rm RoomController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var habitacion dto.HabitacionDTO
	if err := json.NewDecoder(r.Body).Decode(&habitacion); err != nil {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Habitacion{
		Numero: habitacion.Numero,
		Tipo:   habitacion.Tipo,
		Precio: habitacion.Precio,
		Estado: habitacion.Estado,
	}

	findHabitacion := rm.Hb.Create(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findHabitacion)
}

func (rm RoomController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var habitacion dto.HabitacionDTO
	if err := json.NewDecoder(r.Body).Decode(&habitacion); err != nil {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Habitacion{
		ID:     idStr,
		Numero: habitacion.Numero,
		Tipo:   habitacion.Tipo,
		Precio: habitacion.Precio,
		Estado: habitacion.Estado,
	}

	findHabitacion := rm.Hb.Mod(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findHabitacion)
}

func (rm RoomController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener habitacion")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Habitacion{ID: idStr}
	findHabitacion := rm.Hb.Del(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findHabitacion)
}
