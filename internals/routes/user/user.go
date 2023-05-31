package noteRoutes

import (
	userhandler "twitter-public-api/internals/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/:username")
	// Read all Notes
	note.Get("/", userhandler.GetProfile)
	note.Get("/avatar", userhandler.GetAvatar)
	// // Read one Note
	note.Get("/tweets", userhandler.GetTweets)
}
