package repositories

import "github.com/janjos/user-api/entities"

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id int) (*entities.User, error)
	LogIn(email, password string) (*entities.User, error)
}
