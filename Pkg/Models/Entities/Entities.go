package entities

import "time"

// Entidad Usuario
// trabajada
type Usuario struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id_usuario"`
	TipoDocumento   string `gorm:"type:enum('CC','TI','TE','PP','PPT','NIT')" json:"tipo_documento"`
	NumeroDocumento string `gorm:"size:50;unique" json:"numero_documento"`
	Nombre          string `gorm:"size:100" json:"nombre"`
	Apellido        string `gorm:"size:100" json:"apellido"`
	Email           string `gorm:"size:100;unique" json:"email"`
	Telefono        string `gorm:"size:20" json:"telefono"`
	Nacionalidad    string `gorm:"size:100" json:"nacionalidad"`
	Ciudad          string `gorm:"size:100" json:"ciudad"`
	Pais            string `gorm:"size:100" json:"pais"`
	Ocupacion       string `gorm:"size:100" json:"ocupacion"`
	PaisProcedencia string `gorm:"size:100" json:"pais_procedencia"`
	Direccion       string `gorm:"size:255" json:"direccion"`
}

// Entidad Usuarios
type Usuarios []Usuario

// Entidad Habitacion
// trabajada
type Habitacion struct {
	ID     int     `gorm:"primaryKey;autoIncrement" json:"id_habitacion"`
	Numero string  `gorm:"size:10" json:"numero"`
	Tipo   string  `gorm:"type:enum('sencilla','doble','suite')" json:"tipo"`
	Precio float64 `gorm:"type:decimal(10,2)" json:"precio"`
	Estado string  `gorm:"type:enum('disponible','reservada','ocupada','no disponible')" json:"estado"`
}

type Habitaciones []Habitacion

// Entidad Reserva

type Reserva struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id_reserva"`
	FechaReserva time.Time  `json:"fecha_reserva"`
	Estado       string     `gorm:"type:enum('confirmada','cancelada')" json:"estado"`
	IDUsuario    int        `json:"id_usuario"`
	IDHabitacion int        `json:"id_habitacion"`
	Usuario      Usuario    `gorm:"foreignKey:IDUsuario" json:"usuario"`
	Habitacion   Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

type Reservas []Reserva

// Entidad Pago
type Pago struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id_pago"`
	IDReserva  int       `json:"id_reserva"`
	Monto      float64   `gorm:"type:decimal(10,2)" json:"monto"`
	MetodoPago string    `gorm:"type:enum('tarjeta','efectivo')" json:"metodo_pago"`
	FechaPago  time.Time `json:"fecha_pago"`
	Reserva    Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

type Pagos []Pago

// Entidad Llave
type Llave struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id_llave"`
	IDReserva    int        `json:"id_reserva"`
	IDHabitacion int        `json:"id_habitacion"`
	TipoLlave    string     `gorm:"type:enum('fisica','electronica')" json:"tipo_llave"`
	EstadoLlave  string     `gorm:"type:enum('activa','desactivada')" json:"estado_llave"`
	Reserva      Reserva    `gorm:"foreignKey:IDReserva" json:"reserva"`
	Habitacion   Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

type Llaves []Llave

// Entidad CheckInCheckOut
type CheckInCheckOut struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id_checkin"`
	IDReserva     int       `json:"id_reserva"`
	FechaCheckIn  time.Time `json:"fecha_checkin"`
	FechaCheckOut time.Time `json:"fecha_checkout"`
	Reserva       Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

type CheckInCheckOuts []CheckInCheckOut

// Entidad Personal
type Personal struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id_personal"`
	Nombre   string `gorm:"size:100" json:"nombre"`
	Apellido string `gorm:"size:100" json:"apellido"`
	Rol      string `gorm:"type:enum('recepcionista','limpieza','gerente','botones','supervisor')" json:"rol"`
}

type Personales []Personal

// Entidad FacturaElectronica
type FacturaElectronica struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id_factura"`
	IDPago       int       `json:"id_pago"`
	FechaFactura time.Time `json:"fecha_factura"`
	Total        float64   `gorm:"type:decimal(10,2)" json:"total"`
	Pago         Pago      `gorm:"foreignKey:IDPago" json:"pago"`
}

type FacturaElectronicas []FacturaElectronica

// Entidad PersonalHabitacion
type PersonalHabitacion struct {
	ID              int        `gorm:"primaryKey;autoIncrement" json:"id_personal_habitacion"`
	IDPersonal      int        `json:"id_personal"`
	IDHabitacion    int        `json:"id_habitacion"`
	FechaAsignacion time.Time  `json:"fecha_asignacion"`
	Personal        Personal   `gorm:"foreignKey:IDPersonal" json:"personal"`
	Habitacion      Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

type PersonalHabitaciones []PersonalHabitacion
