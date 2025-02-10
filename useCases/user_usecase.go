package useCases

import (
	"github.com/janjos/user-api/entities"
	"github.com/janjos/user-api/external"
	"github.com/janjos/user-api/interfaces/repositories"
)

type UserUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) CreateUser(email, password string) (*entities.User, error) {
	user := &entities.User{
		Email:    email,
		Password: password,
	}

	err := uc.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) GetUserByID(id int) (*entities.User, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) LogIn(email, password string) (*entities.User, error) {
	user, err := uc.userRepo.LogIn(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) VerifyToken(token string) (float64, error) {
	id, err := external.VerifyToken(token)
	if err != nil {
		return -1, err
	}

	return id, nil
}
