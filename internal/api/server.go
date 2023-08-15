package api

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture/internal/api/route"
	"go-clean-architecture/internal/config"
	"golang.org/x/exp/slog"
)

func Start(cfg *config.Config) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	route.NewIndexRoute(app)

	slog.Info("Http server starting at :3000")
	_ = app.Listen(":3000")
}
