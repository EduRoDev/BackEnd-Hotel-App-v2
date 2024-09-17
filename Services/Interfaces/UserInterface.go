package interfaces

type User interface {
	Login()
	SingUp()
}

type UserRoom interface {
	GetRoom()
	GetRooms()
	SelectRoom()
	ModifyRoom()
	DeleteRoom()
}
