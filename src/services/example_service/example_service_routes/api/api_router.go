package auth_service_api_routes

import (
	"github.com/gofiber/fiber/v2"
	v1Routes "golang-fiber-boilerplate-api/src/services/example_service/example_service_routes/api/v1"
	v2Routes "golang-fiber-boilerplate-api/src/services/example_service/example_service_routes/api/v2"
)

func Routes(route fiber.Router) {
	route.Get("/", handler)
	route.Route("/v1", v1Routes.Routes)
	route.Route("/v2", v2Routes.Routes)
}

func handler(c *fiber.Ctx) error {
	return c.JSON(c.Route())
}
