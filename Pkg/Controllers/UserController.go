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
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gorilla/mux"
)

type UserController struct {
	l  *log.Logger
	Us interfaces.User
}

func NewUserController(l *log.Logger) *UserController {
	return &UserController{l, &impl.User{}}
}

var limite = tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	httpError := tollbooth.LimitByRequest(limite, w, r)
	if httpError != nil {
		w.WriteHeader(httpError.StatusCode)
		w.Write([]byte(httpError.Message))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var loginRequest entities.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	user, err := u.Us.Login(loginRequest.Nombre, loginRequest.NumeroDocumento)
	if err != nil {
		rp := helpers.ErrorWithStatus("Error", "Error al conectar con el servidor", "500")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := u.Us.Get()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}

func (u UserController) GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	user := entities.Usuario{ID: idStr}
	findUser := u.Us.GetID(user)

	if findUser.ID == 0 {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findUser)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	nombre := vars["nombre"]
	User := u.Us.GetUser(nombre)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}

func (u UserController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dto.UsuarioDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Usuario{
		TipoDocumento:   user.TipoDocumento,
		NumeroDocumento: user.NumeroDocumento,
		Nombre:          user.Nombre,
		Apellido:        user.Apellido,
		Email:           user.Email,
		Telefono:        user.Telefono,
		Nacionalidad:    user.Nacionalidad,
		Ciudad:          user.Ciudad,
		Pais:            user.Pais,
		Ocupacion:       user.Ocupacion,
		PaisProcedencia: user.PaisProcedencia,
		Direccion:       user.Direccion,
	}

	findUser := u.Us.Create(data)

	var err error
	if findUser["error"] != nil {
		rp := helpers.Error(err, "Error al crear usuario")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(rp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findUser)
}

func (u UserController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var user dto.UsuarioDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Usuario{
		ID:              idStr,
		TipoDocumento:   user.TipoDocumento,
		NumeroDocumento: user.NumeroDocumento,
		Nombre:          user.Nombre,
		Apellido:        user.Apellido,
		Email:           user.Email,
		Telefono:        user.Telefono,
		Nacionalidad:    user.Nacionalidad,
		Ciudad:          user.Ciudad,
		Pais:            user.Pais,
		Ocupacion:       user.Ocupacion,
		PaisProcedencia: user.PaisProcedencia,
		Direccion:       user.Direccion,
	}

	findUser := u.Us.Mod(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findUser)
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := helpers.Error(err, "Error al obtener usuario")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Usuario{ID: idStr}
	findUser := u.Us.Del(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findUser)
}
