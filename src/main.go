package main

import (
	"log/slog"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// MiddleWare
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/compile/swagger.yaml",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		slog.Error("Failed to start Fiber Server")
	}
}
