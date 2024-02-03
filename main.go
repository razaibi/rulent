package main

import (
	"log"
	"sync"

	"rulent/handlers"
	"rulent/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := models.Config{}
	yamlFile := "events.yaml"
	config.ParseYAML(yamlFile)

	errorChan := make(chan error)
	var wg sync.WaitGroup

	app.Post("/validate", handlers.ValidateHandler(&config, errorChan, &wg))
	// Goroutine for handling errors from the asynchronous actions
	go func() {
		for err := range errorChan {
			log.Printf("error from action: %v", err)
			// Handle the error, e.g., by logging or sending a notification
		}
	}()

	app.Get("/reload", handlers.ReloadHandler(&config, yamlFile))
	log.Fatal(app.Listen(":8081"))

	// Close the error channel when all goroutines are done
	wg.Wait()
	close(errorChan)
}
