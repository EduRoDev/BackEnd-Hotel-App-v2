package entities

import "time"

type Check struct {
	Id        int
	ReserID   int
	Reser     Reservation
	DateCheck time.Time
	DateOut   time.Time
}

type Checks []Check
