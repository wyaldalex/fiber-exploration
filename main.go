package main

import (
	"github.com/gofiber/fiber/v2"
)

/*
   go mod init FiberAuth
   go get github.com/gofiber/fiber/v2
*/

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Initial Project Setup")
	})

	app.Listen(":10001")
}
