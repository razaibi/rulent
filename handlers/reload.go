package handlers

import (
	"rulent/models"

	"github.com/gofiber/fiber/v2"
)

func ReloadHandler(config *models.Config, yamlFile string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := config.ReloadYAML(yamlFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to reload conditions",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Conditions reloaded successfully",
		})
	}
}
