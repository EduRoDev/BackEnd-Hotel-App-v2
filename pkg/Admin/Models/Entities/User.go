package entities

type User struct {
	Id       int
	Name     string
	LastName string
	Email    string
	Password string
	Addres   string
	Phone    string
}

type Users []User
