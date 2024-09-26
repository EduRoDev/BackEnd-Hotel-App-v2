package dtos

import "time"


// DTO Usuario
type UsuarioDTO struct {
	ID        uint   `json:"id_usuario"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	Telefono  string `json:"telefono"`
	Direccion string `json:"direccion"`
}

// DTO Habitacion
type HabitacionDTO struct {
	ID     uint    `json:"id_habitacion"`
	Numero string  `json:"numero"`
	Tipo   string  `json:"tipo"`
	Precio float64 `json:"precio"`
	Estado string  `json:"estado"`
}

// DTO Reserva
type ReservaDTO struct {
	ID           uint      `json:"id_reserva"`
	IDUsuario    uint      `json:"id_usuario"`
	IDHabitacion uint      `json:"id_habitacion"`
	FechaReserva time.Time `json:"fecha_reserva"`
	Estado       string    `json:"estado"`
}

// DTO Pago
type PagoDTO struct {
	ID         uint      `json:"id_pago"`
	IDReserva  uint      `json:"id_reserva"`
	Monto      float64   `json:"monto"`
	MetodoPago string    `json:"metodo_pago"`
	FechaPago  time.Time `json:"fecha_pago"`
}

// DTO Llave
type LlaveDTO struct {
	ID           uint   `json:"id_llave"`
	IDReserva    uint   `json:"id_reserva"`
	IDHabitacion uint   `json:"id_habitacion"`
	TipoLlave    string `json:"tipo_llave"`
	EstadoLlave  string `json:"estado_llave"`
}

// DTO CheckInCheckOut
type CheckInCheckOutDTO struct {
	ID            uint      `json:"id_checkin"`
	IDReserva     uint      `json:"id_reserva"`
	FechaCheckIn  time.Time `json:"fecha_checkin"`
	FechaCheckOut time.Time `json:"fecha_checkout"`
}

// DTO Personal
type PersonalDTO struct {
	ID       uint   `json:"id_personal"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Rol      string `json:"rol"`
}

// DTO FacturaElectronica
type FacturaElectronicaDTO struct {
	ID           uint      `json:"id_factura"`
	IDPago       uint      `json:"id_pago"`
	FechaFactura time.Time `json:"fecha_factura"`
	Total        float64   `json:"total"`
}

// DTO PersonalHabitacion
type PersonalHabitacionDTO struct {
	ID              uint      `json:"id_personal_habitacion"`
	IDPersonal      uint      `json:"id_personal"`
	IDHabitacion    uint      `json:"id_habitacion"`
	FechaAsignacion time.Time `json:"fecha_asignacion"`
}
