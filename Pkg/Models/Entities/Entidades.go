package entities

import "time"

// Entidad Usuario
type Usuario struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id_usuario"`
	Nombre    string `gorm:"size:100" json:"nombre"`
	Apellido  string `gorm:"size:100" json:"apellido"`
	Email     string `gorm:"size:100;unique" json:"email"`
	Telefono  string `gorm:"size:20" json:"telefono"`
	Direccion string `gorm:"size:255" json:"direccion"`
}

// Entidad Habitacion
type Habitacion struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id_habitacion"`
	Numero string  `gorm:"size:10" json:"numero"`
	Tipo   string  `gorm:"type:enum('sencilla','doble','suite')" json:"tipo"`
	Precio float64 `gorm:"type:decimal(10,2)" json:"precio"`
	Estado string  `gorm:"type:enum('disponible','reservada','ocupada','no disponible')" json:"estado"`
}

// Entidad Reserva
type Reserva struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id_reserva"`
	IDUsuario    uint       `json:"id_usuario"`
	IDHabitacion uint       `json:"id_habitacion"`
	FechaReserva time.Time  `json:"fecha_reserva"`
	Estado       string     `gorm:"type:enum('confirmada','cancelada')" json:"estado"`
	Usuario      Usuario    `gorm:"foreignKey:IDUsuario" json:"usuario"`
	Habitacion   Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

// Entidad Pago
type Pago struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id_pago"`
	IDReserva  uint      `json:"id_reserva"`
	Monto      float64   `gorm:"type:decimal(10,2)" json:"monto"`
	MetodoPago string    `gorm:"type:enum('tarjeta','efectivo')" json:"metodo_pago"`
	FechaPago  time.Time `json:"fecha_pago"`
	Reserva    Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

// Entidad Llave
type Llave struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id_llave"`
	IDReserva    uint       `json:"id_reserva"`
	IDHabitacion uint       `json:"id_habitacion"`
	TipoLlave    string     `gorm:"type:enum('fisica','electronica')" json:"tipo_llave"`
	EstadoLlave  string     `gorm:"type:enum('activa','desactivada')" json:"estado_llave"`
	Reserva      Reserva    `gorm:"foreignKey:IDReserva" json:"reserva"`
	Habitacion   Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

// Entidad CheckInCheckOut
type CheckInCheckOut struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id_checkin"`
	IDReserva     uint      `json:"id_reserva"`
	FechaCheckIn  time.Time `json:"fecha_checkin"`
	FechaCheckOut time.Time `json:"fecha_checkout"`
	Reserva       Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

// Entidad Personal
type Personal struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id_personal"`
	Nombre   string `gorm:"size:100" json:"nombre"`
	Apellido string `gorm:"size:100" json:"apellido"`
	Rol      string `gorm:"type:enum('recepcionista','limpiesa','gerente','botones','supervisor')" json:"rol"`
}

// Entidad FacturaElectronica
type FacturaElectronica struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id_factura"`
	IDPago       uint      `json:"id_pago"`
	FechaFactura time.Time `json:"fecha_factura"`
	Total        float64   `gorm:"type:decimal(10,2)" json:"total"`
	Pago         Pago      `gorm:"foreignKey:IDPago" json:"pago"`
}

// Entidad PersonalHabitacion
type PersonalHabitacion struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id_personal_habitacion"`
	IDPersonal      uint       `json:"id_personal"`
	IDHabitacion    uint       `json:"id_habitacion"`
	FechaAsignacion time.Time  `json:"fecha_asignacion"`
	Personal        Personal   `gorm:"foreignKey:IDPersonal" json:"personal"`
	Habitacion      Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}
