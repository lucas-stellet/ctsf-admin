package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-stellet/ctsf-admin/server/application/usecases"
	"github.com/lucas-stellet/ctsf-admin/server/domain"
	"github.com/lucas-stellet/ctsf-admin/server/framework/utils"
	"net/http"
)

type Controller interface {
	Handle(ctx *fiber.Ctx)
}

type CreateUserController struct {
	CreateUserUseCase usecases.CreateUserUseCase
}

func (c CreateUserController) Handle(ctx *fiber.Ctx) error {
	body := new(domain.User)

	if err := utils.ParseBody(ctx, body); err != nil {
		return err
	}

	user, err := domain.NewUser(body.Name, body.Email, body.Password)

	if err != nil {
		return err
	}

	_ , err = c.CreateUserUseCase.Execute(user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"info": "user created",
	})
}
