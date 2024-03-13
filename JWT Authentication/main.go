package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pushpa/database"
	"github.com/pushpa/routes"
)


func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
	}))

   	routes.Routes(app)

    app.Listen(":8080")
}