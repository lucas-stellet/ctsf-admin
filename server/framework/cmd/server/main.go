package main

import (
	"log"
	"path"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lucas-stellet/ctsf-admin/server/application/routes"
	"github.com/lucas-stellet/ctsf-admin/server/framework/utils"

	"os"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../../../")

	err := godotenv.Load(d + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env files: %v", err)
	}
}

func main() {
	db := utils.ConnectDB(os.Getenv("env"))

	app := fiber.New()

	userRouter := app.Group("/users")

	userRouter.Post("/", routes.CreateUserRoute(db))

	app.Listen(":3000")
}
