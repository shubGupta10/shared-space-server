package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shubGupta10/shared-space-server/internals/config"
	"github.com/shubGupta10/shared-space-server/internals/models"
	"github.com/shubGupta10/shared-space-server/internals/utils"
)

func CreateSpace(c *fiber.Ctx) error {
	var data map[string]string

	//parse the request body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Body"})
	}

	//validate the input data
	Token := data["token"]
	Creator := data["creator"]
	Partner := data["partner"]
	Name := data["name"]

	if Token == "" || Creator == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please enter Token and Creator"})
	}

	//Creator and Partner should be UUIDs

	spaceItem := models.Space{
		ID:        uuid.New(),
		Name:      Name,
		Token:     Token,
		Creator:   utils.ConvertToUUID(Creator),
		Partner:   utils.ConvertToUUID(Partner),
		CreatedAt: time.Now(),
	}

	//create the space in the database
	if err := config.DB.Create(&spaceItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create space"})
	}

	return c.JSON(spaceItem)
}

func DeleteSpace(c *fiber.Ctx) error {
	var data map[string]string

	//parse the request body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "INvalid body"})
	}

	//validate the spaceId
	spaceId := data["spaceId"]
	if spaceId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please enter spaceId"})
	}

	//delete the space from the database
	if err := config.DB.Where("id = ?", spaceId).Delete(&models.Space{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete space"})
	}

	return c.JSON(fiber.Map{"message": "Space deleted successfully"})
}

func FetchSpace(c *fiber.Ctx) error {
	var spaces []models.Space

	//get the current userid
	userId := c.Locals("user_id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UserId not found"})
	}

	//convert the userId to UUID
	newUUID, err := uuid.Parse(userId.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	//fetch all spaces where user is either creator and partner
	if err := config.DB.Where("creator = ? OR partner = ?", newUUID, newUUID).Find(&spaces).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch the space"})
	}

	return c.JSON(fiber.Map{"message": "Successfully fetched the spaces", "spaces": spaces})
}
