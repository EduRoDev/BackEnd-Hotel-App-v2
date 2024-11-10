package interfaces

import (
	"time"

	entities "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Models/Entities"
)

type User interface {
	Login(email string, numero_documento string) (string, error)
	Get() []entities.Usuario
	GetID(User entities.Usuario) entities.Usuario
	GetUser(nombre string, apellido string) entities.Usuario
	Create(User entities.Usuario) map[string]interface{}
	Mod(User entities.Usuario) map[string]interface{}
	Del(User entities.Usuario) map[string]interface{}
}

type Acompañante interface {
	Get() []entities.Acompañante
	GetID(Acompañante entities.Acompañante) entities.Acompañante
	Create(Acompañante entities.Acompañante) map[string]interface{}
	Mod(Acompañante entities.Acompañante) map[string]interface{}
	Del(Acompañante entities.Acompañante) map[string]interface{}
}

type Habitacion interface {
	Get() []entities.Habitacion
	GetID(Habitacion entities.Habitacion) entities.Habitacion
	GetAvailable() []entities.Habitacion
	Create(Habitacion entities.Habitacion) map[string]interface{}
	Mod(Habitacion entities.Habitacion) map[string]interface{}
	Del(Habitacion entities.Habitacion) map[string]interface{}
}

type Reservation interface {
	Get() []entities.Reserva
	GetID(Reserva entities.Reserva) entities.Reserva
	GetByUsuarioYFecha(idUsuario int, fechaEntrada time.Time) entities.Reserva
	Create(Reserva entities.Reserva) map[string]interface{}
	Mod(Reserva entities.Reserva) map[string]interface{}
	Del(Reserva entities.Reserva) map[string]interface{}
}

type Payment interface {
	Get() []entities.Pago
	GetID(Pago entities.Pago) entities.Pago
	GetByIdReserva(idReserva int) entities.Pago
	Create(Pago entities.Pago) map[string]interface{}
	Mod(Pago entities.Pago) map[string]interface{}
	Del(Pago entities.Pago) map[string]interface{}
	Cancel(idReserva int) map[string]interface{}
}

type Personal interface {
	Get() []entities.Personal
	GetID(Personal entities.Personal) entities.Personal
	Asing(Personal entities.Personal) map[string]interface{}
	Mod(Personal entities.Personal) map[string]interface{}
	Del(Personal entities.Personal) map[string]interface{}
}

type PersonalHabitacion interface {
	Get() []entities.PersonalHabitacion
	GetID(PersonalHabitacion entities.PersonalHabitacion) entities.PersonalHabitacion
	Asing(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
	Mod(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
	Del(PersonalHabitacion entities.PersonalHabitacion) map[string]interface{}
}
