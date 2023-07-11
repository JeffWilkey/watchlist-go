package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWatchlists(c *fiber.Ctx) error {
	// Get ownerId from query parameters
	ownerId, err := primitive.ObjectIDFromHex(c.Query("ownerId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing ownerId"})
	}

	// Get all watchlists from database
	results, err := service.FindWatchlistsByOwnerId(c, ownerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get watchlists", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All watchlists", "data": dto.CreateWatchlistsResponse(results)})
}

func CreateWatchlist(c *fiber.Ctx) error {
	watchlist := new(model.Watchlist)

	if err := c.BodyParser(watchlist); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "data": err})
	}

	// Validate watchlist input
	validate := validator.New()
	err := validate.Struct(watchlist)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "data": err})
	}

	// Create watchlist in database
	err = service.CreateWatchlist(c, watchlist)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create watchlist", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created watchlist", "data": dto.CreateWatchlistResponse(*watchlist)})
}

func UpdateWatchlist(c *fiber.Ctx) error {
	var input dto.WatchlistUpdateRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Validate watchlist input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Format watchlist id from param and get token
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid watchlist id", "data": err})
	}

	token := c.Locals("user").(*jwt.Token)
	watchlist, err := service.FindWatchlistById(c, id)

	// Validate user is owner of watchlist
	if !service.ValidToken(token, watchlist.OwnerID) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	// Update watchlist in database
	status, err := service.UpdateWatchlist(c, id, input, watchlist)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated watchlist", "data": dto.CreateWatchlistResponse(*watchlist)})
}

func DeleteWatchlist(c *fiber.Ctx) error {
	return c.SendString("Delete watchlist")
}
