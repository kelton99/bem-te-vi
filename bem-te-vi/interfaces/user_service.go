package interfaces

import "application/domain"

type UserService interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(id int) (domain.User, error)
	GetUserByEmail(username string) (domain.User, error)
	GetUsers() ([]domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUser(id int) error
}
