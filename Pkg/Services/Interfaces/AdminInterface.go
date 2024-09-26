package interfaces

type Admin interface {
	Login()
}

type AdminStrategy interface {
	Get()
	GetAll()
	Asig()
	Mod()
	Del()
}


