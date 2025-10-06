package rest

import (
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func RegisterMiddleware(app *fiber.App) {

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/compile/swagger.yaml",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(swagger.New(cfg))

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))
}
