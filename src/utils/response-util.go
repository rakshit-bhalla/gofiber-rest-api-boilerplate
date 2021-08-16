package utils

import (
	"rakshit.dev/gofiber-rest-api-boilerplate/src/models"
)

type Response = models.Response

func GetResponse(data interface{}, err error) Response {
	if err != nil {
		return Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
	}
	return Response{
		Success: true,
		Data:    data,
		Error:   "",
	}
}
