package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shankar7042/hotel-reservation-golang/types"
)

func HandleGetUsers(c fiber.Ctx) error {
	users := []types.User{
		{
			FirstName: "Shankar",
			LastName:  "Kumar",
		},
		{
			FirstName: "Ashok",
			LastName:  "Kumar",
		},
		{
			FirstName: "Santosh",
			LastName:  "Kumar",
		},
	}
	return c.JSON(users)
}

func HandleGetUser(c fiber.Ctx) error {
	return c.JSON(types.User{
		FirstName: "Shankar",
		LastName:  "Kumar",
	})
}
