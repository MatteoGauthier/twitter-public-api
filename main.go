package main

import (
	"log"
	"twitter-public-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	log.Fatal(app.Listen(":3000"))

}
