package pages

import (
	"fmt"

	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandlePost(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'self'; img-src 'self'; font-src 'self'; block-all-mixed-content")

	post, err := api.FetchPosts(c.Params("postID"))
	if post.Id == "" {
		post, err = api.FetchPosts(c.Params("postID"))
		if post.Id == "" {
			post, err = api.FetchMedia(c.Params("postID"))
			if post.Id == "" {
				return fmt.Errorf("404 page not found")				
			}
		}
	}
 	if err != nil {
		return err
	}

	comments := []types.Comment{}
	if post.SharedWithCommunity {
		c.Set("Cache-Control", "public,max-age=604800")
		comments, err = api.FetchComments(c.Params("postID"))
		if err != nil {
			return err
		}
	} else {
		c.Set("Cache-Control", "public,max-age=31557600")
	}

	return c.Render("post", fiber.Map{
		"post":    post,
		"comments": comments,
	})
}