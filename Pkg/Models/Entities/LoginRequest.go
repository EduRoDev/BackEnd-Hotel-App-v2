package entities

type LoginResponse struct {
	Token   string  `json:"token"`
	Reserva Reserva `json:"reserva"`
}

type LoginRequest struct {
	Email           string `json:"email"`
	NumeroDocumento string `json:"numero_documento"`
}
