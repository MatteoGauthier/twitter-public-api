package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"

	userRoutes "twitter-public-api/internals/routes/user"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	userRoutes.SetupNoteRoutes(api)
}
