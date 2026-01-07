package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handles all errors in a consistent format
func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   err.Error(),
		"status":  code,
		"path":    c.Path(),
		"method":  c.Method(),
	})
}
