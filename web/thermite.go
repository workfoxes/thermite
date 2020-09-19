package web

import (
	"github.com/airstrik/thermite/pkg/server"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// StartThermite : will start the application server with certain parameter
func StartThermite(name string, port int) {
	_server := server.CreateAppServer(name, port)

	// Add Middleware to Application
	_server.Use(compress.New(compress.Config{Level: compress.LevelBestCompression}))

	_server.Start()
}
