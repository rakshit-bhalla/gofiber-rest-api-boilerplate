package routers

import (
	"github.com/gofiber/fiber/v2"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/controllers"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/middlewares"
)

type UserController = controllers.UserController

// SetUserRoutes ...
func SetUserRoutes(router *fiber.App, userController UserController) {

	userRouter := router.Group("/users")
	userRouter.Get("/", userController.GetUsers)
	userRouter.Post("/", middlewares.ValidateCreateUserReq, userController.CreateUser)
	userRouter.Get("/:userId", userController.GetUser)
	userRouter.Delete("/:userId", userController.DeleteUser)

}
