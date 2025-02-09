package repositories

import (
	"context"

	"github.com/janjos/user-api/entities"
	"github.com/janjos/user-api/external"
)

type UserRepositoryImpl struct {
	db *external.DbConnection
}

func NewUserRepositoryImpl(db *external.DbConnection) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Save(user *entities.User) error {
	return r.db.Db.QueryRow(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password", user.Email, user.Password).Scan(&user.Id, &user.Email, &user.Password)
}

func (r *UserRepositoryImpl) FindByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.Db.QueryRow(context.Background(), "SELECT id, email FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
