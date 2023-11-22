package models

type User struct {
	Id    int
	Name  string
	Email string
}

type UserUsecase interface {
	CreateUser(user *User) (int, error)
	GetUserByID(id int) (*User, error)
	GetAllUsers() ([]*User, error)
}

type UserRepository interface {
	Create(user *User) (int, error)
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
}
