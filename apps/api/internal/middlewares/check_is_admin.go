package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/gigaservices/libs/models"
)

func IsAdmin(c *fiber.Ctx) error {
	raw_user := c.Locals("user")

	if raw_user == nil {
		return fiber.NewError(401, "unauthentificated")
	}

	user := raw_user.(models.User)

	if user.IsAdmin {
		return c.Next()
	}

	return c.JSON(fiber.Map{"message": "you are not an admin"})
}
