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

type PaymentController struct {
	l  *log.Logger
	Py interfaces.Payment
}

func NewPaymentController(l *log.Logger) *PaymentController {
	return &PaymentController{l, &impl.Pago{}}
}

func (py PaymentController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pago := py.Py.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(pago)
}

func (py PaymentController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	pago := entities.Pago{ID: idStr}
	findPago := py.Py.GetID(pago)
	if findPago.ID == 0 {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPago)
}

func (py PaymentController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pago dto.PagoDTO
	if err := json.NewDecoder(r.Body).Decode(&pago); err != nil {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Pago{
		ID:         pago.ID,
		IDReserva:  pago.IDReserva,
		Monto:      pago.Monto,
		MetodoPago: pago.MetodoPago,
		FechaPago:  pago.FechaPago,
	}

	findPago := py.Py.Create(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findPago)
}

func (py PaymentController) Mod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var pago entities.Pago
	if err := json.NewDecoder(r.Body).Decode(&pago); err != nil {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Pago{
		ID:         idStr,
		IDReserva:  pago.IDReserva,
		Monto:      pago.Monto,
		MetodoPago: pago.MetodoPago,
		FechaPago:  pago.FechaPago,
	}

	findPago := py.Py.Mod(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPago)
}

func (py PaymentController) Del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener pago")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Pago{ID: idStr}
	findPago := py.Py.Del(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findPago)
}

func (py PaymentController) Cancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}