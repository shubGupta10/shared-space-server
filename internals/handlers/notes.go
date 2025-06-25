package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubGupta10/shared-space-server/internals/config"
	"github.com/shubGupta10/shared-space-server/internals/models"
	"github.com/shubGupta10/shared-space-server/internals/utils"
)

func CreateNote(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Body"})
	}

	//validate the input data
	spaceId := data["spaceId"]
	content := data["content"]
	author := data["author"]

	if spaceId == "" || content == "" || author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please enter spaceId, content and author"})
	}

	//create the note in db
	noteItem := models.Notes{
		SpaceID: utils.ConvertToUUID(spaceId),
		Author:  utils.ConvertToUUID(author),
		Content: content,
	}

	//save the note in db
	if err := config.DB.Create(&noteItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create note"})
	}

	return c.JSON(noteItem)
}

func DeleteNote(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Body"})
	}

	//validate the noteId
	noteId := data["noteId"]
	if noteId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please enter noteId"})
	}

	//delete the note from the database
	if err := config.DB.Where("id = ?", noteId).Delete(&models.Notes{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete note"})
	}

	return c.JSON(fiber.Map{"message": "Note deleted successfully"})
}
