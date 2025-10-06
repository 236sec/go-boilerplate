package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/rest"
)

func main() {
	app := fiber.New()

	rest.RegisterMiddleware(app)

	rest.RouteRegisterHandlers(app)

	err := app.Listen(":3000")
	if err != nil {
		slog.Error("Failed to start Fiber Server")
	}
}
