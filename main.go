package main

import (
	"log"

	"github.com/jeffwilkey/what-to-watch/database"
	"github.com/jeffwilkey/what-to-watch/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDb()

	router.SetupRoutes(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	});

	log.Fatalln(app.Listen(":3000"))
}