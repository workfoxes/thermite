package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ApplicationServer : Application server will hold the service object for the application
type ApplicationServer struct {
	fiber.App
	Name string
	Port int
}

// CreateAppServer : func to create Application server object to Manage the application server
func CreateAppServer(Name string, Port int) *ApplicationServer {
	_server := &ApplicationServer{Name: Name, Port: Port}
	return _server
}

func (app *ApplicationServer) Start() {
	_port := strconv.Itoa(app.Port)
	log.Fatal(app.Listen(":" + _port))
}
