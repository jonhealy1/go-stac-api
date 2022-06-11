package main

import (
	"go-stac-api/configs"
	"go-stac-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	routes.CollectionRoute(app)

	app.Listen(":6000")
}
