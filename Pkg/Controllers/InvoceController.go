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

type InvoiceController struct {
	l  *log.Logger
	Il interfaces.Factura
}

func NewInvoiceController(l *log.Logger) *InvoiceController {
	return &InvoiceController{l, &impl.Invoice{}}
}

func (il InvoiceController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	invoice := il.Il.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(invoice)
}

func (il InvoiceController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	invoice := entities.FacturaElectronica{ID: idStr}
	findInvoice := il.Il.GetID(invoice)
	if findInvoice.ID == 0 {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findInvoice)
}

func (il InvoiceController) Asign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var invoice dto.FacturaElectronicaDTO
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.FacturaElectronica{
		ID:           invoice.ID,
		IDPago:       invoice.IDPago,
		FechaFactura: invoice.FechaFactura,
		Total:        invoice.Total,
	}
	findInvoice := il.Il.Asing(data)
	if findInvoice["status"] == "error" {
		rp := helpers.ErrorWithStatus("Error in database", "Error al crear factura", "error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findInvoice)
}

func (il InvoiceController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var invoice dto.FacturaElectronicaDTO
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.FacturaElectronica{
		ID:           idStr,
		IDPago:       invoice.IDPago,
		FechaFactura: invoice.FechaFactura,
		Total:        invoice.Total,
	}

	findInvoice := il.Il.Mod(data)

	if findInvoice["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findInvoice)
}

func (il InvoiceController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.FacturaElectronica{ID: idStr}
	findInvoice := il.Il.Del(data)

	if findInvoice["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener factura")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findInvoice)
}
