package entities

import "time"

type Reservation struct {
	Id      int
	UserID  int
	User    User
	RoomID  int
	Room    Room
	DateIn  time.Time
	State   []string
}

type Reservations []Reservation