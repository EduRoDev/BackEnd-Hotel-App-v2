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

type CheckController struct {
	l  *log.Logger
	Cl interfaces.CheckInCheckOut
}

func NewCheckController(l *log.Logger) *CheckController {
	return &CheckController{l, &impl.CheckInCheckOut{}}
}

func (ck CheckController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	checkInCheckOut := ck.Cl.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(checkInCheckOut)
}

func (ck CheckController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	checkInCheckOut := entities.CheckInCheckOut{ID: idStr}
	findCheckInCheckOut := ck.Cl.GetID(checkInCheckOut)
	if findCheckInCheckOut.ID == 0 {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findCheckInCheckOut)
}

func (ck CheckController) Asign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var checkInCheckOut dto.CheckInCheckOutDTO
	if err := json.NewDecoder(r.Body).Decode(&checkInCheckOut); err != nil {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.CheckInCheckOut{
		ID:            checkInCheckOut.ID,
		IDReserva:     checkInCheckOut.IDReserva,
		FechaCheckIn:  checkInCheckOut.FechaCheckIn,
		FechaCheckOut: checkInCheckOut.FechaCheckOut,
	}
	findCheckInCheckOut := ck.Cl.Asing(data)
	if findCheckInCheckOut["status"] == "error" {
		rp := helpers.ErrorWithStatus("Error in database", "Error al crear checkInCheckOut", "error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findCheckInCheckOut)
}

func (ck CheckController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var checkInCheckOut dto.CheckInCheckOutDTO
	if err := json.NewDecoder(r.Body).Decode(&checkInCheckOut); err != nil {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.CheckInCheckOut{
		ID:            idStr,
		FechaCheckIn:  checkInCheckOut.FechaCheckIn,
		FechaCheckOut: checkInCheckOut.FechaCheckOut,
	}

	findCheckInCheckOut := ck.Cl.Mod(data)

	if findCheckInCheckOut["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findCheckInCheckOut)
}

func (ck CheckController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.CheckInCheckOut{ID: idStr}
	findCheckInCheckOut := ck.Cl.Del(data)

	if findCheckInCheckOut["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener checkInCheckOut")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findCheckInCheckOut)
}
