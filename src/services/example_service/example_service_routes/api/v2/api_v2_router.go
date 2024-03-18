package api_v2_router

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber-boilerplate-api/src/services/example_service/example_service_routes/api/v2/api_v2_example_router"
)

func Routes(route fiber.Router) {
	route.Get("/", handler)
	route.Route("/examples", api_v2_example_router.Routes)
}

func handler(c *fiber.Ctx) error {
	return c.JSON(c.Route())
}
