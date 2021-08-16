package main

import (
	"log"

	"rakshit.dev/gofiber-rest-api-boilerplate/src/configs"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/controllers"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/repositories"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/routers"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.SetUserRoutes(app, userController)

	log.Fatal(app.Listen(configs.APIHost))
}
