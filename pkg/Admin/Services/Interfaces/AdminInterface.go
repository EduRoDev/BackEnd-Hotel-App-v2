package interfaces

type AdminRoom interface {
	GetRoom()
	GetRooms()
	AsigRoom()
	ModRoom()
	DelRoom()
}

type AdminPersonal interface {
	GetPerson()
	GetPersons()
	AsigPerson()
	ModPerson()
	DelPerson()
}
