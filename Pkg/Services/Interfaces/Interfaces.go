package interfaces

import entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"

type User interface {
	Get() []entities.Usuario
	GetID(User entities.Usuario) entities.Usuario
	Asing(User entities.Usuario) map[string]interface{}
	Mod(User entities.Usuario) map[string]interface{}
	Del(User entities.Usuario) map[string]interface{}
}

type Habitation interface {
	Get() []entities.Habitacion
	GetID(Habitacion entities.Habitacion) entities.Habitacion
	Asing(Habitacion entities.Habitacion) map[string]interface{}
	Mod(Habitacion entities.Habitacion) map[string]interface{}
	Del(Habitacion entities.Habitacion) map[string]interface{}
}

type Reservation interface {
	Get() []entities.Reserva
	GetID(Reserva entities.Reserva) entities.Reserva
	Asing(Reserva entities.Reserva) map[string]interface{}
	Mod(Reserva entities.Reserva) map[string]interface{}
	Del(Reserva entities.Reserva) map[string]interface{}
}

type Payment interface {
	Get() []entities.Pago
	GetID(Pago entities.Pago) entities.Pago
	Asing(Pago entities.Pago) map[string]interface{}
	Mod(Pago entities.Pago) map[string]interface{}
	Del(Pago entities.Pago) map[string]interface{}
}

type Key interface {
	Get() []entities.Llave
	GetID(Llave entities.Llave) entities.Llave
	Asing(Llave entities.Llave) map[string]interface{}
	Mod(Llave entities.Llave) map[string]interface{}
	Del(Llave entities.Llave) map[string]interface{}
}

type CheckInCheckOut interface {
	Get() []entities.CheckInCheckOut
	GetID(CheckInCheckOut entities.CheckInCheckOut) entities.CheckInCheckOut
	Asing(CheckInCheckOut entities.CheckInCheckOut) map[string]interface{}
	Mod(CheckInCheckOut entities.CheckInCheckOut) map[string]interface{}
	Del(CheckInCheckOut entities.CheckInCheckOut) map[string]interface{}
}

type Personal interface {
	Get() []entities.Personal
	GetID(Personal entities.Personal) entities.Personal
	Asing(Personal entities.Personal) map[string]interface{}
	Mod(Personal entities.Personal) map[string]interface{}
	Del(Personal entities.Personal) map[string]interface{}
}

type Factura interface {
	Get() []entities.FacturaElectronica
	GetID(FacturaElectronica entities.FacturaElectronica) entities.FacturaElectronica
	Asing(FacturaElectronica entities.FacturaElectronica) map[string]interface{}
	Mod(FacturaElectronica entities.FacturaElectronica) map[string]interface{}
	Del(FacturaElectronica entities.FacturaElectronica) map[string]interface{}
}

type PersonalHabitacion interface {
	Get() []entities.PersonalHabitacion
	GetID(PersonalHabitacion entities.PersonalHabitacion) entities.PersonalHabitacion
	Asing(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
	Mod(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
	Del(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
}

type Admin interface {
	Get() []entities.Usuario
	GetID(User entities.Usuario) entities.Usuario
	Asing(User entities.Usuario) map[string]interface{}
	Mod(User entities.Usuario) map[string]interface{}
	Del(User entities.Usuario) map[string]interface{}
}
