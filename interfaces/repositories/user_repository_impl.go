package repositories

import (
	"context"
	"fmt"

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
	hashedPassword, _ := external.HashPassword(user.Password)
	return r.db.Db.QueryRow(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password", user.Email, hashedPassword).Scan(&user.Id, &user.Email, &user.Password)
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

func (r *UserRepositoryImpl) LogIn(email, password string) (*entities.User, error) {
	var user entities.User
	err := r.db.Db.QueryRow(context.Background(), "SELECT id, email, password FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	match := external.VerifyPassword(password, user.Password)

	if !match {
		fmt.Println("Error logging in user - hash verify = $1", match)
		return nil, err
	}

	if user.Token != "" {
		fmt.Println("User already logged in")
		return &user, nil
	}

	tokenString, err := external.CreateToken(user.Email, user.Id)
	if err != nil {
		fmt.Errorf("No email found")
		return nil, err
	}
	user.Token = tokenString

	fmt.Println("User logged in! Token:$1 ", user.Token)

	return &user, nil
}
