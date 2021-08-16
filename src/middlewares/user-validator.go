package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/models"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/utils"
)

type UserReq = models.UserReq

func ValidateCreateUserReq(c *fiber.Ctx) error {
	var input UserReq
	if err := c.BodyParser(&input); err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusBadRequest)
		c.JSON(utils.GetResponse(nil, errors.New("bad params")))
		return nil
	}
	if input.Email == "" {
		c.Status(http.StatusBadRequest)
		c.JSON(utils.GetResponse(nil, errors.New("bad params")))
		return nil
	}
	return c.Next()
}
