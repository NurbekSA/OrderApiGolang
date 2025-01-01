package exception

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Handle(ex *AppError, c *fiber.Ctx) error {
	switch ex.Code {
	case 400:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": ex.Error()})
	case 500:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": ex.Error()})
	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": ex.Error()})
	}
}
