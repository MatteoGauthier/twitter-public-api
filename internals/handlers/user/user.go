package userHandler

import (
	"context"
	"fmt"
	"strings"
	logging "twitter-public-api/internals/utils"

	"github.com/gofiber/fiber/v2"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

var scraper = twitterscraper.New()

func GetProfile(c *fiber.Ctx) error {
	profile, err := scraper.GetProfile(c.Params("username"))
	if err != nil {
		// log
		logging.ErrorLog(fmt.Sprint(err))
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// @todo Add avatar large image
	return c.JSON(profile)
}

func GetAvatar(c *fiber.Ctx) error {
	profile, err := scraper.GetProfile(c.Params("username"))
	if err != nil {
		panic(err)
	}

	largeImage := strings.Replace(profile.Avatar, "normal", "400x400", 1)
	return c.Redirect(largeImage)
}

func GetTweets(c *fiber.Ctx) error {
	tweets := scraper.GetTweets(context.Background(), c.Params("username"), 10)
	// @todo Transform tweets to a more readable format
	return c.JSON(tweets)
}
