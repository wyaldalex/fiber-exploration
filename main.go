package main

import (
	"FiberAuth/database"
	"FiberAuth/routes"

	"github.com/gofiber/fiber/v2"
)

/*
   go mod init FiberAuth
   go get github.com/gofiber/fiber/v2
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
*/

func main() {

	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":10001")
}
