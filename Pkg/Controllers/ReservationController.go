package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

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

func (rs ReservationController) GetByUsuarioYFecha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idUsuario, err := strconv.Atoi(vr["idUsuario"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}
	fechaEntrada, err := time.Parse("2006-01-02", vr["fechaEntrada"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener reserva")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}
	reserva := rs.Rs.GetByUsuarioYFecha(idUsuario, fechaEntrada)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(reserva)
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

	fechaReserva, _ := time.Parse("2006-01-02 15:04:05", reserva.FechaReserva+" 00:00:00")
	fechaEntrada, _ := time.Parse("2006-01-02 15:04:05", reserva.FechaEntrada+" 00:00:00")
	fechaSalida, _ := time.Parse("2006-01-02 15:04:05", reserva.FechaSalida+" 00:00:00")

	data := entities.Reserva{
		ID:           reserva.ID,
		FechaReserva: fechaReserva,
		FechaEntrada: fechaEntrada,
		FechaSalida:  fechaSalida,
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
		FechaEntrada: reserva.FechaEntrada,
		FechaSalida:  reserva.FechaSalida,
		Estado:       reserva.Estado,
	}

	findReserva := rs.Rs.Mod(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findReserva)
}


func (rs ReservationController) Del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener el ID de la reserva desde los parámetros de la URL
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener ID de la reserva")
		w.WriteHeader(http.StatusBadRequest) // Cambiado a 400 Bad Request
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Reserva{ID: idStr}
	resultado := rs.Rs.Del(data)

	// Verificar si hubo un error en la operación
	if resultado["status"] == "error" {
		w.WriteHeader(http.StatusInternalServerError) // Error en el servidor si la operación falló
		json.NewEncoder(w).Encode(resultado)
		return
	}

	// Enviar respuesta de éxito si no hubo errores
	w.WriteHeader(http.StatusAccepted) // 202 Accepted indica que la operación fue aceptada
	json.NewEncoder(w).Encode(resultado)
}
