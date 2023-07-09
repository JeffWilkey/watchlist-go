package handler

import "github.com/gofiber/fiber/v2"

func GetWatchlists(c *fiber.Ctx) error {
	return c.SendString("All watchlists")
}

func GetWatchlist(c *fiber.Ctx) error {
	return c.SendString("Single watchlist")
}

func CreateWatchlist(c *fiber.Ctx) error {
	return c.SendString("Create watchlist")
}

func UpdateWatchlist(c *fiber.Ctx) error {
	return c.SendString("Update watchlist")
}

func DeleteWatchlist(c *fiber.Ctx) error {
	return c.SendString("Delete watchlist")
}
