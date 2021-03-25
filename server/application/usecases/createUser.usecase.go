package usecases

import (
	"github.com/lucas-stellet/ctsf-admin/server/application/repositories"
	"github.com/lucas-stellet/ctsf-admin/server/domain"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *CreateUserUseCase) Execute(user *domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
