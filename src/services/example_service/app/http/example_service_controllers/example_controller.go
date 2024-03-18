package example_service_controllers

import "github.com/gofiber/fiber/v2"

func Example(c *fiber.Ctx) error {
	return c.JSON(c.Route())
}
