package api

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/shankar7042/hotel-reservation-golang/db"
	"github.com/shankar7042/hotel-reservation-golang/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c fiber.Ctx) error {
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
