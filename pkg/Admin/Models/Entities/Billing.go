package entities

import "time"

type Billing struct {
	Id          int
	PayId       int
	Pay         Pay
	DateBilling time.Time
	Quantity     float64
}

type Billings []Billing
