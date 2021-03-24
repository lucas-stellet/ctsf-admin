package usecases

import (
	"fmt"

	"github.com/lucas-stellet/ctsf-admin/server/application/repositories"
	"github.com/lucas-stellet/ctsf-admin/server/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) Auth(email, password string) (*domain.User, error) {
	user, err := u.UserRepository.Find(email)

	if err != nil {
		return nil, fmt.Errorf("the password is invalid for the user: %v", email)
	}

	if user.IsCorrectPassword(password) {
		return user, nil
	}

	return nil, fmt.Errorf("the password is invalid for the user: %v", email)
}
