package pages

import (
	"os"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleFrontpage(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=31557600")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'; style-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	return c.Render("frontpage", fiber.Map{
		"proto": c.Protocol(),
		"domain": c.Hostname(),
		"force_webp": os.Getenv("FORCE_WEBP"),
	})
}