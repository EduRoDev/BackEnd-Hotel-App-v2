package entities

import "time"

// Entidad Usuario
type Usuario struct {
	ID                 int          `gorm:"primaryKey;autoIncrement" json:"id_usuario"`
	TipoDocumento      string       `gorm:"type:enum('CC','TI','TE','PP','PPT','NIT')" json:"tipo_documento"`
	NumeroDocumento    string       `gorm:"size:50;unique" json:"numero_documento"`
	Nombre             string       `gorm:"size:100" json:"nombre"`
	Apellido           string       `gorm:"size:100" json:"apellido"`
	Email              string       `gorm:"size:100;unique" json:"email"`
	Edad               int          `gorm:"type:int" json:"edad"`
	Telefono           string       `gorm:"size:20" json:"telefono"`
	Ciudad             string       `gorm:"size:100" json:"ciudad"`
	Pais               string       `gorm:"size:100" json:"pais"`
	Direccion          string       `gorm:"size:255" json:"direccion"`
	Acompañante        Acompañantes `gorm:"foreignKey:IDusuario" json:"acompañantes"`
}

type Usuarios []Usuario

// Definir TableName para Usuario
func (Usuario) TableName() string {
	return "usuario"
}

type Acompañante struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id_acompañante"`
	IDusuario       int    `gorm:"index" json:"usuario"`
	Nombre          string `gorm:"size:100" json:"nombre"`
	Apellido        string `gorm:"size:100" json:"apellido"`
	TipoDocumento   string `gorm:"type:enum('CC','TI','TE','PP','PPT','NIT')" json:"tipo_documento"`
	NumeroDocumento string `gorm:"size:50;unique" json:"numero_documento"`
	NumeroTelefono  string `gorm:"size:50" json:"numero_telefono"`
}

type Acompañantes []Acompañante

// Definir TableName para Acompañante
func (Acompañante) TableName() string {
	return "acompañante"
}

// Entidad Habitacion
type Habitacion struct {
	ID     int     `gorm:"primaryKey;autoIncrement" json:"id_habitacion"`
	Nombre string  `gorm:"size:100" json:"nombre"`
	Tipo   string  `gorm:"type:enum('sencilla','doble','suite')" json:"tipo"`
	Precio float64 `gorm:"type:decimal(10,2)" json:"precio"`
	Estado string  `gorm:"type:enum('disponible','reservada','ocupada')" json:"estado"`
}

type Habitaciones []Habitacion

// Definir TableName para Habitacion
func (Habitacion) TableName() string {
	return "habitacion"
}

// Entidad Reserva
type Reserva struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id_reserva"`
	FechaReserva time.Time  `json:"fecha_reserva"`
	FechaEntrada time.Time  `json:"fecha_entrada"`
	FechaSalida  time.Time  `json:"fecha_salida"`
	Estado       string     `gorm:"type:enum('confirmada','reservada','pendiente','cancelada')" json:"estado"`
	IDUsuario    int        `json:"id_usuario"`
	IDHabitacion int        `json:"id_habitacion"`
	Usuario      Usuario    `gorm:"foreignKey:IDUsuario" json:"usuario"`
	Habitacion   Habitacion `gorm:"foreignKey:IDHabitacion" json:"habitacion"`
}

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
	Estado     string    `gorm:"type:enum('pendiente','cancelada','realizado')" json:"estado"`
	FechaPago  time.Time `json:"fecha_pago"`
	Reserva    Reserva   `gorm:"foreignKey:IDReserva" json:"reserva"`
}

type Pagos []Pago

// Definir TableName para Pago
func (Pago) TableName() string {
	return "pago"
}

// Entidad Personal
type Personal struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id_personal"`
	Nombre   string `gorm:"size:100" json:"nombre"`
	Apellido string `gorm:"size:100" json:"apellido"`
	Rol      string `gorm:"type:enum('recepcionista','limpieza','gerente','botones','supervisor')" json:"rol"`
}

type Personals []Personal

// Definir TableName para Personal
func (Personal) TableName() string {
	return "personal"
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

type PersonalHabitacions []PersonalHabitacion

// Definir TableName para PersonalHabitacion
func (PersonalHabitacion) TableName() string {
	return "personal_habitacion"
}
