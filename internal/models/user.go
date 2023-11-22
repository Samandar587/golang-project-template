package models

type User struct {
	Id    int
	Name  string
	Email string
}

type UserRepository interface {
	Create(user User) (int, error)
	GetByID(id int) (User, error)
	GetAll() ([]User, error)
}
