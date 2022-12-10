package middlewares

import (
	"github.com/devhammed/fibery/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

var validate = validator.New()

func Validator(value interface{}) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(value); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&models.JsonError{
				Ok:      false,
				Message: "Unable to parse request.",
			})
		}

		var errors []*ErrorResponse

		err := validate.Struct(value)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, &ErrorResponse{
					Field:   err.Field(),
					Error:   err.Tag(),
					Message: err.Error(),
				})
			}
		}

		if errors != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
		}

		c.Locals("validatedBody", value)

		return c.Next()
	}
}
