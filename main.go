package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//start a new fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API is up!")
		return err
	})

	//port no 3000
	app.Listen(":3000")
}
