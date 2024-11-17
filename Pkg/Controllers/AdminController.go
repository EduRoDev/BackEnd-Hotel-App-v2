package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	auth "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Auth"
	helpers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Helpers"
	dto "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Dto"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
	impl "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Impl"
	interfaces "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Interfaces"
	"github.com/didip/tollbooth"
)

type AdminController struct {
	logger *log.Logger
	admin  interfaces.Administrador
}

func NewAdminController(logger *log.Logger) *AdminController {
	return &AdminController{logger, &impl.Admin{}}
}

func (am *AdminController) Login(w http.ResponseWriter, r *http.Request) {
	httpError := tollbooth.LimitByRequest(limite, w, r)
	if httpError != nil {
		w.WriteHeader(httpError.StatusCode)
		w.Write([]byte(httpError.Message))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var loginRequest entities.LoginRequestAdmin
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		rp := helpers.Error(err, "Error al obtener el request")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	passwordHash := auth.HashPassword(loginRequest.Password)
	token, err := am.admin.Login(loginRequest.Email, passwordHash)
	if err != nil {
		rp := helpers.ErrorWithStatus("Error", "Error al conectar con el servidor", "500")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(token)
}

func (am *AdminController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var admin dto.AdministradorDTO
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		rp := helpers.Error(err, "Error al obtener el request")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Administrador{
		Nombre:   admin.Nombre,
		Apellido: admin.Apellido,
		Email:    admin.Email,
		Password: auth.HashPassword(admin.Password),
	}

	findAdmin := am.admin.Create(data)

	if findAdmin["error"] != nil {
		rp := helpers.Error(nil, "Error al crear el admin")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findAdmin)
}
