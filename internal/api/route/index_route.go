package route

import "github.com/gofiber/fiber/v2"

func NewIndexRoute(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"hello": "world"})
	})
}
