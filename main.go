package main

import (
	"FiberAuth/database"
	"FiberAuth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*
   go mod init FiberAuth
   go get github.com/gofiber/fiber/v2
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   github.com/dgrijalva/jwt-go/v4
*/

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	//app.Listen(":10001")
	app.Listen("127.0.0.1:10001") // add localhost to avoid firewall warning in development
}
