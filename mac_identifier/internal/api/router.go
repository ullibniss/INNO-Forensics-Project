package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func NewRouter(macHandler *MacHandler) *fiber.App {
	app := fiber.New()

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./resources"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./resources/html/index.html")
	})

	api := app.Group("")
	macHandler.addMacRouter(api)

	return app
}
