package entities

import "time"

type Plan struct {
	Id          int
	Description string
	Date        time.Time
}

type Plans []Plan
