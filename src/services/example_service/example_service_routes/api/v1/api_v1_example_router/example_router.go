package api_v1_example_router

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber-boilerplate-api/src/services/example_service/app/http/example_service_controllers"
)

func Routes(route fiber.Router) {
	//group := route.Group("/v1")

	route.Get("/", example_service_controllers.Example)
	route.Get("/:id", example_service_controllers.Example)
	route.Post("/", example_service_controllers.Example)
	route.Put("/:id", example_service_controllers.Example)
	route.Delete("/:id", example_service_controllers.Example)
}
