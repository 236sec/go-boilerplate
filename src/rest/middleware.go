package rest

import (
	"log"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"goboilerplate.com/src/rest/response"
)

func RegisterMiddleware(app *fiber.App) {
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/compile/swagger.yaml",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	app.Use(cors.New())

	app.Use(swagger.New(cfg))

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	app.Use(RecoveryMiddleware)
}

func RecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v\n", r)

			res := response.Responses[response.InternalServerErrorResponse]
			if err := c.Status(res.HttpStatus).JSON(res); err != nil {
				log.Printf("Failed to write recovery response: %v\n", err)
			}
		}
	}()
	return c.Next()
}