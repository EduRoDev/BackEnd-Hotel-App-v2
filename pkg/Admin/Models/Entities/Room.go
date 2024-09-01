package entities

type Room struct {
	Id        int
	NumRoom   string
	TypeRoom  []string
	PriceRoom float64
	State     []string
}

type Rooms []Room
