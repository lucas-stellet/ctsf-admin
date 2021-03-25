package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-stellet/ctsf-admin/server/application/controllers"
	"github.com/lucas-stellet/ctsf-admin/server/application/repositories"
	"github.com/lucas-stellet/ctsf-admin/server/application/usecases"
	"gorm.io/gorm"
)



func CreateUserRoute(db *gorm.DB) func(ctx *fiber.Ctx) error {
	userRepository := repositories.UserRepositoryDb{Db: db}

	createUseCase := usecases.CreateUserUseCase{
		UserRepository: userRepository,
	}

	createUserController := controllers.CreateUserController{CreateUserUseCase: createUseCase}

	return func(ctx *fiber.Ctx) error {
		return createUserController.Handle(ctx)
	}
}
