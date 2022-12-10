package controllers

import (
	"github.com/devhammed/fibery/database"
	"github.com/devhammed/fibery/models"
	"github.com/devhammed/fibery/requests"
	"github.com/gofiber/fiber/v2"
)

func ApiV1_GetUsers(c *fiber.Ctx) error {
	return c.JSON(&models.JsonResponse{
		Ok:   true,
		Data: database.GetUsers(),
	})
}

func ApiV1_GetUser(c *fiber.Ctx) error {
	return c.JSON(&models.JsonResponse{
		Ok: true,
		Data: models.User{
			ID:   1,
			Name: "Hammed Oyedele",
		},
	})
}

func ApiV1_CreateUser(c *fiber.Ctx) error {
	data := c.Locals("validatedBody").(*requests.CreateUserRequest)

	database.InsertUser(&models.User{
		ID:   len(database.GetUsers()) + 1,
		Name: data.Name,
	})

	return c.JSON(&models.JsonResponse{
		Ok:      true,
		Message: "User added.",
	})
}
