package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from DeployAnchor Backend!")
	})

	fmt.Println("Starting server on port 3000")
	app.Listen(":3000")
}
