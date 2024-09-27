package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	dto "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Dto"
	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
	impl "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Impl"
	interfaces "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Services/Interfaces"
	"github.com/gorilla/mux"
)

type UserController struct {
	l  *log.Logger
	Us interfaces.User
}

func NewUserController(l *log.Logger) *UserController {
	return &UserController{l, &impl.User{}}
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
		rp := map[string]interface{}{
			"error":   err,
			"message": "Error al obtener usuario"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}
	user := entities.Usuario{ID: idStr}
	findUser := u.Us.GetID(user)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findUser)
}

func (u UserController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dto.UsuarioDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rp := map[string]string{
			"error":   err.Error(),
			"message": "Error al obtener usuario",
		}
		w.WriteHeader(http.StatusNotFound)
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

	findUser := u.Us.Asing(data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(findUser)
}

func (u UserController) Modify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		rp := map[string]interface{}{
			"error":   err,
			"message": "Error al obtener usuario"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	var user dto.UsuarioDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rp := map[string]interface{}{"error": err, "message": "Error al obtener usuario"}
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
		rp := map[string]interface{}{
			"error":   err,
			"message": "Error al obtener usuario"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(rp)
		return
	}

	data := entities.Usuario{ID: idStr}
	findUser := u.Us.Del(data)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(findUser)
}
