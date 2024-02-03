package handlers

import (
	"rulent/logic"
	"rulent/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func ValidateHandler(config *models.Config, errorChan chan<- error, wg *sync.WaitGroup) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload map[string]interface{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		// Validate JSON against events and rules here
		outcomes, isValid := logic.ValidateJSON(payload, config)
		if isValid {
			// Execute actions based on the mode specified in each outcome
			logic.ExecuteActions(outcomes, payload, errorChan, wg)

			return c.JSON(fiber.Map{
				"status":   "success",
				"outcomes": outcomes,
			})
		} else {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Validation failed",
			})
		}
	}
}
