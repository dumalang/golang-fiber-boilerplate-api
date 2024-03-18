package example_service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-fiber-boilerplate-api/src/services/example_service/example_service_routes"
)

func Index() {
	fmt.Println("example_service.Main")
	App()
}

func App() {
	fmt.Println("example_service.App")
	app := fiber.New()
	app.Route("/example", example_service_routes.Routes)
	app.Get("", base)
	app.Listen(":3001")
}

func base(c *fiber.Ctx) error {
	return c.SendString("golang-fiber-boilerplate-api")
}
