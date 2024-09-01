package entities

import "time"

type Pay struct {
	Id            int
	ReservationId int
	Reser         Reservation
	Quantity      float64
	MethodPay     []string
	DatePay       time.Time
}

type Pays []Pay
