package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
   go mod init FiberAuth
   go get github.com/gofiber/fiber/v2
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
*/

func main() {

	_, err := gorm.Open(mysql.Open("root:root@/fiberauthdb"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Initial Project Setup")
	})

	app.Listen(":10001")
}
