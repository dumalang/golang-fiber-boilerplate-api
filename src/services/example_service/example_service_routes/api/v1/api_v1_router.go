package auth_service_api_v1_routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber-boilerplate-api/src/services/example_service/example_service_routes/api/v1/api_v1_example_router"
)

func Routes(route fiber.Router) {
	route.Get("/", handler)
	route.Route("/examples", api_v1_example_router.Routes)
}

func handler(c *fiber.Ctx) error {
	return c.SendString("example-service.api.v1")
}
