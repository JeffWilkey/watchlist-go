package main

import (
	"log"

	"github.com/jeffwilkey/watchlist-go/config"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/router"

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