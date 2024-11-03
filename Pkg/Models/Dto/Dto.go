package dto

import "time"

// DTO para Usuario
type UsuarioDTO struct {
	ID              int              `json:"id_usuario"`
	TipoDocumento   string           `json:"tipo_documento"`
	NumeroDocumento string           `json:"numero_documento"`
	Nombre          string           `json:"nombre"`
	Apellido        string           `json:"apellido"`
	Email           string           `json:"email"`
	FechaNacimiento string        `json:"fecha_nacimiento"`
	Telefono        string           `json:"telefono"`
	Ciudad          string           `json:"ciudad"`
	Pais            string           `json:"pais"`
	Direccion       string           `json:"direccion"`
	Acompañantes    []AcompañanteDTO `json:"acompañantes"`
}

// DTO para Acompañante
type AcompañanteDTO struct {
	ID              int    `json:"id_acompañante"`
	Nombre          string `json:"nombre"`
	TipoDocumento   string `json:"tipo_documento"`
	NumeroDocumento string `json:"numero_documento"`
}

// DTO para Habitacion
type HabitacionDTO struct {
	ID     int     `json:"id_habitacion"`
	Nombre string  `json:"nombre"`
	Tipo   string  `json:"tipo"`
	Precio float64 `json:"precio"`
	Estado string  `json:"estado"`
}

// DTO para Reserva
type ReservaDTO struct {
	ID           int       `json:"id_reserva"`
	FechaReserva time.Time `json:"fecha_reserva"`
	FechaEntrada time.Time `json:"fecha_entrada"`
	FechaSalida  time.Time `json:"fecha_salida"`
	Estado       string    `json:"estado"`
	IDUsuario    int       `json:"id_usuario"`
	IDHabitacion int       `json:"id_habitacion"`
}

// DTO para Pago
type PagoDTO struct {
	ID         int       `json:"id_pago"`
	IDReserva  int       `json:"id_reserva"`
	Monto      float64   `json:"monto"`
	MetodoPago string    `json:"metodo_pago"`
	FechaPago  time.Time `json:"fecha_pago"`
}

// DTO para Personal
type PersonalDTO struct {
	ID       int    `json:"id_personal"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Rol      string `json:"rol"`
}

// DTO para PersonalHabitacion
type PersonalHabitacionDTO struct {
	ID              int       `json:"id_personal_habitacion"`
	IDPersonal      int       `json:"id_personal"`
	IDHabitacion    int       `json:"id_habitacion"`
	FechaAsignacion time.Time `json:"fecha_asignacion"`
}
