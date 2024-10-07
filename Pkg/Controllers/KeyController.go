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

type KeyController struct {
	l  *log.Logger
	Kl interfaces.Key
}

func NewKeyController(l *log.Logger) *KeyController {
	return &KeyController{l, &impl.Llave{}}
}

func (kl KeyController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	llave := kl.Kl.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(llave)
}

func (kl KeyController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	llave := entities.Llave{ID: idStr}
	findLlave := kl.Kl.GetID(llave)
	if findLlave.ID == 0 {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findLlave)
}

func (kl KeyController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var llave dto.LlaveDTO
	if err := json.NewDecoder(r.Body).Decode(&llave); err != nil {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	data := entities.Llave{
		ID:          llave.ID,
		IDReserva:   llave.IDReserva,
		TipoLlave:   llave.TipoLlave,
		EstadoLlave: llave.EstadoLlave,
	}
	findLlave := kl.Kl.Create(data)
	if findLlave["status"] == "error" {
		rp := helpers.ErrorWithStatus("Error in database", "Error al crear llave", "error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findLlave)
}

func (kl KeyController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var llave dto.LlaveDTO
	if err := json.NewDecoder(r.Body).Decode(&llave); err != nil {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Llave{
		ID:          idStr,
		TipoLlave:   llave.TipoLlave,
		EstadoLlave: llave.EstadoLlave,
	}

	findLlave := kl.Kl.Mod(data)

	if findLlave["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findLlave)
}

func (kl KeyController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Llave{ID: idStr}
	findLlave := kl.Kl.Del(data)

	if findLlave["status"] == "error" {
		rp := helpers.Error(err, "Error al obtener llave")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findLlave)
}
