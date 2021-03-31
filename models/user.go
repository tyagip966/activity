package models

type User struct {
	Id    int
	Name  string
	Email string
	Phone string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
	CreateUser(user User) (int, error)
	UpdateUser(user User) (int, error)
	DeleteUser(id int) (int, error)
}