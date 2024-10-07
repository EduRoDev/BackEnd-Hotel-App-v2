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

type PersonalRoomController struct {
	l  *log.Logger
	Pl interfaces.PersonalHabitacion
}

func NewPersonalRoomController(l *log.Logger) *PersonalRoomController {
	return &PersonalRoomController{l, &impl.PersonalRoom{}}
}

func (pl PersonalRoomController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personalRoom := pl.Pl.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(personalRoom)
}

func (pl PersonalRoomController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	personalRoom := entities.PersonalHabitacion{ID: idStr}
	findPersonalRoom := pl.Pl.GetID(personalRoom)
	if findPersonalRoom.ID == 0 {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonalRoom)
}

func (pl PersonalRoomController) Asign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var personalRoom dto.PersonalHabitacionDTO
	if err := json.NewDecoder(r.Body).Decode(&personalRoom); err != nil {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.PersonalHabitacion{
		ID:              personalRoom.ID,
		IDPersonal:      personalRoom.IDPersonal,
		IDHabitacion:    personalRoom.IDHabitacion,
		FechaAsignacion: personalRoom.FechaAsignacion,
	}
	findPersonalRoom := pl.Pl.Asing(data)
	if findPersonalRoom["status"] == "error" {
		rp := helpers.ErrorWithStatus("Error in database", "Error al crear personalRoom", "error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findPersonalRoom)
}

func (pl PersonalRoomController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var personalRoom dto.PersonalHabitacionDTO
	if err := json.NewDecoder(r.Body).Decode(&personalRoom); err != nil {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.PersonalHabitacion{
		ID:              idStr,
		IDPersonal:      personalRoom.IDPersonal,
		IDHabitacion:    personalRoom.IDHabitacion,
		FechaAsignacion: personalRoom.FechaAsignacion,
	}

	findPersonalRoom := pl.Pl.Mod(data)

	if findPersonalRoom["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonalRoom)
}

func (pl PersonalRoomController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.PersonalHabitacion{ID: idStr}
	findPersonalRoom := pl.Pl.Del(data)

	if findPersonalRoom["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener personalRoom")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPersonalRoom)
}
