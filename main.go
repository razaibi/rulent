package main

import (
	"flag"
	"log"
	"sync"

	"rulent/handlers"
	"rulent/models"

	"github.com/gofiber/fiber/v2"
)

func main() {

	configPathPtr := flag.String("config", "events.yaml", "path to config file")
	flag.Parse() // Parse the flags

	// Dereference the pointer to get the actual config path
	configFilePath := *configPathPtr

	app := fiber.New(fiber.Config{
		ServerHeader: "Rulent Server",
	})
	config := models.Config{}
	err := config.ParseConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

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

	app.Get("/reload", handlers.ReloadHandler(&config, configFilePath))
	log.Fatal(app.Listen(":8081"))

	// Close the error channel when all goroutines are done
	wg.Wait()
	close(errorChan)
}
