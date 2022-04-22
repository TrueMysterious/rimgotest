package main

import (
	"fmt"
	"net/http"

	"codeberg.org/video-prize-ranch/rimgo/pages"
	"codeberg.org/video-prize-ranch/rimgo/static"
	"codeberg.org/video-prize-ranch/rimgo/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/handlebars"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/etc/rimgu/")
	viper.AddConfigPath("$HOME/.config/rimgu")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("RIMGU_PORT", "3000")
	viper.SetDefault("RIMGU_HOST", "localhost")
	viper.SetDefault("RIMGU_ADDRESS", "0.0.0.0")
	viper.SetDefault("RIMGU_IMGUR_CLIENT_ID", "546c25a59c58ad7")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	engine := handlebars.NewFileSystem(http.FS(views.GetFiles()), ".hbs")
	app := fiber.New(fiber.Config{
		Views:             engine,
		Prefork:           viper.GetBool("FIBER_PREFORK"),
		UnescapePath:      true,
		StreamRequestBody: true,
	})

	app.Use(recover.New())
	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(static.GetFiles()),
	}))

	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		file, _ := static.GetFiles().ReadFile("robots.txt")
		_, err := c.Write(file)
		return err
	})
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		file, _ := static.GetFiles().ReadFile("favicon/favicon.ico")
		_, err := c.Write(file)
		return err
	})

	app.Get("/", pages.HandleFrontpage)
	app.Get("/:postID.gifv", pages.HandleEmbed)
	app.Get("/:baseName.:extension", pages.HandleMedia)
	app.Get("/:postID", pages.HandlePost)
	app.Get("/a/:galleryID", pages.HandleGallery)
	app.Get("/t/:tag", pages.HandleTag)
	app.Get("/user/:userID", pages.HandleUser)
	app.Get("/r/:sub/:postID", pages.HandlePost)
	app.Get("/user/:userID/cover", pages.HandleUserCover)
	app.Get("/user/:userID/avatar", pages.HandleUserAvatar)
	app.Get("/gallery/:galleryID", pages.HandleGallery)

	app.Listen(":" + viper.GetString("RIMGU_PORT"))
}
