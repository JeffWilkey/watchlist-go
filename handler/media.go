package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMediaList(c *fiber.Ctx) error {
	watchlistIdQuery := c.Query("watchlistId")
	watchlistId, err := primitive.ObjectIDFromHex(watchlistIdQuery)
	if watchlistIdQuery == "" || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing or invalid watchlistId"})
	}

	media, err := service.FindMediaByWatchlistId(c, watchlistId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get media", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All media", "data": dto.CreateMediaListResponse(media)})
}

func CreateMedia(c *fiber.Ctx) error {
	media := new(model.Media)

	if err := c.BodyParser(&media); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Validate media input
	validate := validator.New()
	err := validate.Struct(media)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	err = service.CreateMedia(c, media)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create media", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created media", "data": dto.CreateMediaResponse(*media)})
}

func UpdateMedia(c *fiber.Ctx) error {
	var input dto.MediaUpdateRequest
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	watchlistId, err := primitive.ObjectIDFromHex(c.Params("watchlistId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid media ID", "data": err})
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Validate media input
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	token := c.Locals("user").(*jwt.Token)
	watchlist, err := service.FindWatchlistById(c, watchlistId)
	media, err := service.FindMediaById(c, id)

	if !service.ValidToken(token, watchlist.OwnerID) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	// Update media
	status, err := service.UpdateMedia(c, id, input, media)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated watchlist", "data": dto.CreateWatchlistResponse(*watchlist)})
}

func DeleteMedia(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	watchlistId, err := primitive.ObjectIDFromHex(c.Params("watchlistId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid watchlist id", "data": err})
	}

	watchlist, err := service.FindWatchlistById(c, watchlistId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Watchlist not found", "data": err})
	}

	token := c.Locals("user").(*jwt.Token)
	if !service.ValidToken(token, watchlist.OwnerID) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	status, err := service.DeleteMedia(c, id)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted media", "data": nil})
}
