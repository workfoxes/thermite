package web

import (
	"github.com/gofiber/fiber/v2"

	"github.com/workfoxes/thermite/pkg/server"
	"github.com/workfoxes/thermite/web/routes"
)

// StartThermite : Will start the application server with certain parameter
func StartThermite(name string, port int) {
	server.InitDatabaseConnection()

	_server := server.CreateAppServer(name, port)
	_server.LoadDefaultMiddleware()

	_server.Server.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Status": "success",
		})
	})
	api := _server.Server.Group("/api")
	admin := api.Group("/admin")

	routes.AddAdminRoutes(admin)

	_server.Start()
}
