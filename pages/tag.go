package pages

import (
	"strconv"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleTag(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'; style-src 'unsafe-inline' 'self'; media-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	page := "1"
	if c.Query("page") != "" {
		page = c.Query("page")
	}

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 0
	}

	displayPrevPage := true
	if page == "1" {
		displayPrevPage = false
	}

	tag, err := ApiClient.FetchTag(c.Params("tag"), c.Query("sort"), page)
	if err != nil && err.Error() == "ratelimited by imgur" {
		return c.Status(429).Render("errors/429", nil)
	}
	if err != nil {
		return err
	}
	if tag.Display == "" {
		return c.Status(404).Render("errors/404", nil)
	}

	return c.Render("tag", fiber.Map{
		"tag":         tag,
		"page":        page,
		"displayPrev": displayPrevPage,
		"nextPage":    pageNumber + 1,
		"prevPage":    pageNumber - 1,
	})
}
