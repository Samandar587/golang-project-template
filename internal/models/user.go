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
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}

type UserRepository interface {
	Create(user *User) (int, error)
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Update(user *User) (*User, error)
	Delete(id int) error
}
