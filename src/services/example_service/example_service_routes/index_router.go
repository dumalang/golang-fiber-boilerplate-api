package example_service_routes

import (
	"github.com/gofiber/fiber/v2"
	apiRoutes "golang-fiber-boilerplate-api/src/services/example_service/example_service_routes/api"
)

func Routes(route fiber.Router) {
	route.Get("/", handler)
	route.Route("/api", apiRoutes.Routes)
}

func handler(c *fiber.Ctx) error {
	return c.SendString("example-service")
}
