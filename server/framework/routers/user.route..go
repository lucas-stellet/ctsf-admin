package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-stellet/ctsf-admin/server/application/usecases"
	"github.com/lucas-stellet/ctsf-admin/server/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRouters struct {
	UserUseCase usecases.UserUseCase
}

func NewUserRouter() *UserServer {
	return &UserRouters{}
}

func (UserRouters *UserRouters) CreateUser(c *fiber.Ctx) error {

	user, err := domain.NewUser()

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error during the validation: %v", err)
	}

	newUser, err := UserRouters.UserUseCase.Create(user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting information: %v", err)
	}

	return &pb.UserResponse{
		Token: newUser.Token,
	}, nil
}
