package rest

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/di"
	"goboilerplate.com/src/rest/handlers"
)

func RouteRegisterHandlers(app *fiber.App) {
	registerHealthRoutes(app)
}

func registerHealthRoutes(app *fiber.App) {
	healthHandler := handlers.NewHealthHandler(di.GetHealthService())
	app.Get("/health", healthHandler.CheckHealth)
}
