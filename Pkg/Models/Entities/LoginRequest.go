package entities

import esquemas "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Esquemas"

type LoginResponse struct {
	Token   string  `json:"token"`
	Reserva Reserva `json:"reserva"`
}

type LoginRequest struct {
	Email           string `json:"email"`
	NumeroDocumento string `json:"numero_documento"`
}

type loginResponseAdmin struct {
	Token         string                 `json:"token"`
	Administrador esquemas.Administrador `json:"administrador"`
}

type LoginRequestAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
