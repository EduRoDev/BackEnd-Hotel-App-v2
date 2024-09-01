package entities

type Key struct {
	Id int
	ReservationId int
	Reservation Reservation
	RoomId int
	Room Room
	TypeKey []string
	StateKey []bool
}

type Keys []Key
