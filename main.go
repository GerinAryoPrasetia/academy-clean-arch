package main

import (
	"user-management-2/config"
	"user-management-2/handler"
	"user-management-2/repository"
	"user-management-2/routes"
	"user-management-2/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHanlder := handler.NewUserHandler(userUsecase)

	routes.Routes(app, userHanlder)

	app.Listen(":5555")
}
