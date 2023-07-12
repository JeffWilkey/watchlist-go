package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	media, err := service.GetMediaByWatchlistId(c, watchlistId)
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
