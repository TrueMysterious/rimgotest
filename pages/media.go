package pages

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"strings"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleMedia(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=31557600")
	if strings.HasPrefix(c.Path(), "/stack") {
		return handleMedia(c, "https://i.stack.imgur.com/" + strings.ReplaceAll(c.Params("baseName"), "stack/", "") + "." + c.Params("extension"))
	} else {
		return handleMedia(c, "https://i.imgur.com/" + c.Params("baseName") + "." + c.Params("extension"))
	}
}

func HandleUserCover(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=604800")
	return handleMedia(c, "https://imgur.com/user/" + c.Params("userID")  + "/cover?maxwidth=2560")
};

func HandleUserAvatar(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=604800")
	return handleMedia(c, "https://imgur.com/user/" + c.Params("userID")  + "/avatar")
};

func handleMedia(c *fiber.Ctx, url string) error {
	utils.SetHeaders(c)

	if os.Getenv("FORCE_WEBP") == "1" && c.Query("no_webp") == "" && c.Accepts("image/webp") == "image/webp" && !strings.HasPrefix(c.Path(), "/stack") {
		url = strings.ReplaceAll(url, ".png", ".webp")
		url = strings.ReplaceAll(url, ".jpg", ".webp")
		url = strings.ReplaceAll(url, ".jpeg", ".webp")
	}

	if strings.HasPrefix(c.Path(), "/stack") && strings.Contains(c.OriginalURL(), "?") {
		url = url + "?" + strings.Split(c.OriginalURL(), "?")[1]
	}

	optionsHash := ""
	if utils.Config.ImageCache {
		hasher := sha256.New()
		hasher.Write([]byte(url))
		optionsHash = hex.EncodeToString(hasher.Sum(nil))

		image, err := os.ReadFile(utils.Config.CacheDir + "/" + optionsHash)
		if err == nil {
			_, err := c.Write(image)
			return err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if c.Get("Range") != "" {
		req.Header.Set("Range", c.Get("Range"))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		c.Status(404)
		return c.Render("errors/404", nil)
	}

	c.Set("Accept-Ranges", "bytes")
	c.Set("Content-Type", res.Header.Get("Content-Type"));
	c.Set("Content-Length", res.Header.Get("Content-Length"))
	if res.Header.Get("Content-Range") != "" {
		c.Set("Content-Range", res.Header.Get("Content-Range"))
	}

	return c.SendStream(res.Body)
}