package handler

import "github.com/gofiber/fiber"

// Config defines the config for middleware.
type Config struct {
	// PreRequest defines a function to run before each request to start handling the application skip this middleware when returned value.
	//
	// Optional. PreRequest: nil
	PreRequest func(c *fiber.Ctx) bool

	// PostRequest defines a function to run after each request to start handling the application skip this middleware when returned value.
	//
	// Optional. PostRequest: nil
	PostRequest func(c *fiber.Ctx) bool
}
