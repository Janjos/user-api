package repositories

import "user-api/entities"

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id int) (*entities.User, error)
}
