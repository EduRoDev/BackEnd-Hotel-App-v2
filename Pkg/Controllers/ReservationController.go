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

type ReservationController struct {
	l  *log.Logger
	Rs interfaces.Reservation
}

func NewReservationController(l *log.Logger) *ReservationController {
	return &ReservationController{l, &impl.Reserva{}}
}

func (rs ReservationController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reserva := rs.Rs.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(reserva)
}

func (rs ReservationController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	reserva := entities.Reserva{ID: idStr}
	findReserva := rs.Rs.GetID(reserva)
	if findReserva.ID == 0 {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findReserva)
}

func (rs ReservationController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reserva dto.ReservaDTO
	if err := json.NewDecoder(r.Body).Decode(&reserva); err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Reserva{
		ID:           reserva.ID,
		FechaReserva: reserva.FechaReserva,
		Estado:       reserva.Estado,
		IDUsuario:    reserva.IDUsuario,
		IDHabitacion: reserva.IDHabitacion,
	}
	findReserva := rs.Rs.Create(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findReserva)
}

func (rs ReservationController) Mod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var reserva entities.Reserva
	if err := json.NewDecoder(r.Body).Decode(&reserva); err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Reserva{
		ID:           idStr,
		IDUsuario:    reserva.IDUsuario,
		IDHabitacion: reserva.IDHabitacion,
		FechaReserva: reserva.FechaReserva,
		Estado:       reserva.Estado,
	}

	findReserva := rs.Rs.Mod(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findReserva)
}

func (rs ReservationController) Del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Reserva{ID: idStr}
	findReserva := rs.Rs.Del(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findReserva)
}
