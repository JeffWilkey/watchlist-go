package main

import (
	"log"

	"github.com/jeffwilkey/what-to-watch/config"
	"github.com/jeffwilkey/what-to-watch/database"
	"github.com/jeffwilkey/what-to-watch/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	
	// Initialize config and database	
	config.Init()
	database.ConnectDb()

	// Setup routes
	router.SetupRoutes(app)

	log.Fatalln(app.Listen(":3000"))
}