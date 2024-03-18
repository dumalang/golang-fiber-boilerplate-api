package api_v1_example_router

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	//group := route.Group("/v1")

	route.Get("/", example)
	route.Get("/:id", example)
	route.Post("/", example)
	route.Put("/:id", example)
	route.Delete("/:id", example)
}

func example(c *fiber.Ctx) error {
	// Handle logic to get all users
	return c.SendString("api/v1/examples")
}
