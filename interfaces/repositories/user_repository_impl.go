package repositories

import (
	"database/sql"

	"github.com/Janjos/user-api/entities"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Save(user *entities.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Email, user.Password)
	return err
}

func (r *UserRepositoryImpl) FindByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
