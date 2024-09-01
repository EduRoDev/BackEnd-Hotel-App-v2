package interfaces

type User interface {
	CreateUser()
	GetUser()
	UpdateUser()
	DeleteUser()
}

type UserRoom interface {
	GetRoom()
	GetRooms()
	AsigRoom()
	ModRoom()
	DelRoom()
}

