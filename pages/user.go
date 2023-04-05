package pages

import (
	"strconv"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleUser(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'; media-src 'self'; style-src 'unsafe-inline' 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	page := "0"
	if c.Query("page") != "" {
		page = c.Query("page")
	}

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 0
	}

	user, err := ApiClient.FetchUser(c.Params("userID"))
	if err != nil && err.Error() == "ratelimited by imgur" {
		return c.Status(429).Render("errors/429", nil)
	}
	if err != nil {
		return err
	}
	if user.Username == "" {
		return c.Status(404).Render("errors/404", nil)
	}

	submissions, err := ApiClient.FetchSubmissions(c.Params("userID"), "newest", page)
	if err != nil && err.Error() == "ratelimited by imgur" {
		c.Status(429)
		return c.Render("errors/429", nil)
	}
	if err != nil {
		return err
	}

	return c.Render("user", fiber.Map{
		"user":        user,
		"submissions": submissions,
		"page":        page,
		"nextPage": 	 pageNumber + 1,
		"prevPage": 	 pageNumber - 1,
	})
}
