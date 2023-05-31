package userRoutes

import (
	"time"
	"twitter-public-api/internals/handlers"
	userhandler "twitter-public-api/internals/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	note := router.Group("/:username")

	note.Get("/", handlers.Cache(10*time.Minute), userhandler.GetProfile)

	note.Get("/avatar", handlers.Cache(120*time.Minute), userhandler.GetAvatar)

	note.Get("/tweets", handlers.Cache(60*time.Minute), userhandler.GetTweets)
}
