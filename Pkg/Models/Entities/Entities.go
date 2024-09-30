package entities

import "time"

// Entidad Usuario
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

// Recibir un slice de Usuario
type Usuarios []Usuario

// Definir TableName para Usuario
func (Usuario) TableName() string {
	return "usuario"
}

// Entidad Habitacion
type Habitacion struct {
	ID     int     `gorm:"primaryKey;autoIncrement" json:"id_habitacion"`
	Numero string  `gorm:"size:10" json:"numero"`
	Tipo   string  `gorm:"type:enum('sencilla','doble','suite')" json:"tipo"`
	Precio float64 `gorm:"type:decimal(10,2)" json:"precio"`
	Estado string  `gorm:"type:enum('disponible','reservada','ocupada','no disponible')" json:"estado"`
}

// Recibir un slice de Habitacion
type Habitaciones []Habitacion

// Definir TableName para Habitacion
func (Habitacion) TableName() string {
	return "habitacion"
}

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

// Recibir un slice de Reserva
type Reservas []Reserva

// Definir TableName para Reserva
func (Reserva) TableName() string {
	return "reserva"
}

// Entidad Pago
type Pago struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id_pago"`
	IDReserva  int       `json:"id_reserva"`
	Monto      float64   `gorm:"type:decimal(10,2)" json:"monto"`
	MetodoPago string    `gorm:"type:enum('tarjeta','efectivo')" json:"metodo_pago"`
	FechaPago  time.Time `json:"fecha_pago"`
	Reserva    Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

// Recibir un slice de Pago
type Pagos []Pago

// Definir TableName para Pago
func (Pago) TableName() string {
	return "pago"
}

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

// Recibir un slice de Llave
type Llaves []Llave

// Definir TableName para Llave
func (Llave) TableName() string {
	return "llave"
}

// Entidad CheckInCheckOut
type CheckInCheckOut struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id_checkin"`
	IDReserva     int       `json:"id_reserva"`
	FechaCheckIn  time.Time `json:"fecha_checkin"`
	FechaCheckOut time.Time `json:"fecha_checkout"`
	Reserva       Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

// Recibir un slice de CheckInCheckOut
type CheckInCheckOuts []CheckInCheckOut

// Definir TableName para CheckInCheckOut
func (CheckInCheckOut) TableName() string {
	return "checkin_checkout"
}

// Entidad Personal
type Personal struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id_personal"`
	Nombre   string `gorm:"size:100" json:"nombre"`
	Apellido string `gorm:"size:100" json:"apellido"`
	Rol      string `gorm:"type:enum('recepcionista','limpieza','gerente','botones','supervisor')" json:"rol"`
}

// Recibir un slice de Personal
type Personals []Personal

// Definir TableName para Personal
func (Personal) TableName() string {
	return "personal"
}

// Entidad FacturaElectronica
type FacturaElectronica struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id_factura"`
	IDPago       int       `json:"id_pago"`
	FechaFactura time.Time `json:"fecha_factura"`
	Total        float64   `gorm:"type:decimal(10,2)" json:"total"`
	Pago         Pago      `gorm:"foreignKey:IDPago" json:"pago"`
}

// Recibir un slice de FacturaElectronica
type FacturaElectronicas []FacturaElectronica

// Definir TableName para FacturaElectronica
func (FacturaElectronica) TableName() string {
	return "factura_electronica"
}

// Entidad PersonalHabitacion
type PersonalHabitacion struct {
	ID              int        `gorm:"primaryKey;autoIncrement" json:"id_personal_habitacion"`
	IDPersonal      int        `json:"id_personal"`
	IDHabitacion    int        `json:"id_habitacion"`
	FechaAsignacion time.Time  `json:"fecha_asignacion"`
	Personal        Personal   `gorm:"foreignKey:IDPersonal" json:"personal"`
	Habitacion      Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

// Recibir un slice de PersonalHabitacion
type PersonalHabitacions []PersonalHabitacion

// Definir TableName para PersonalHabitacion
func (PersonalHabitacion) TableName() string {
	return "personal_habitacion"
}
