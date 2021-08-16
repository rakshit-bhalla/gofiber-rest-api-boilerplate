package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/models"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/services"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/utils"
)

type User = models.User
type UserReq = models.UserReq
type Response = models.Response

// UserController ...
type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

// NewUserController ...
func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	user, err := u.userService.GetUser(userID)
	if nil != err {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusInternalServerError)
		c.JSON(utils.GetResponse(nil, errors.New("internal server error")))
		return nil
	}
	c.Status(http.StatusOK)
	c.JSON(utils.GetResponse(user, nil))
	return nil
}

func (u *userController) GetUsers(c *fiber.Ctx) error {
	users, err := u.userService.GetUsers()
	if nil != err {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusInternalServerError)
		c.JSON(utils.GetResponse(nil, errors.New("internal server error")))
		return nil
	}
	c.Status(http.StatusOK)
	c.JSON(utils.GetResponse(users, nil))
	return nil
}

func (u *userController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	user, err := u.userService.DeleteUser(userID)
	if nil != err {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusInternalServerError)
		c.JSON(utils.GetResponse(nil, errors.New("internal server error")))
		return nil
	}
	c.Status(http.StatusOK)
	c.JSON(utils.GetResponse(user, nil))
	return nil
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	var input UserReq
	c.BodyParser(&input)
	user, err := u.userService.CreateUser(input)
	if nil != err {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusInternalServerError)
		c.JSON(utils.GetResponse(nil, errors.New("internal server error")))
		return nil
	}
	c.Status(http.StatusCreated)
	c.JSON(utils.GetResponse(user, nil))
	return nil
}
