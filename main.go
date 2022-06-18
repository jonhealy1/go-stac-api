package main

import (
	"go-stac-api/configs"
	"go-stac-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "go-stac-api/docs"
)

// @title go-stac-api
// @version 0.0
// @description STAC api written in go with fiber and mongodb
// @contact.name Jonathan Healy
// @contact.email jonathan.d.healy@gmail.com
// @host localhost:6001
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:6001/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "full",
	}))

	//run database
	configs.ConnectDB()

	routes.CollectionRoute(app)
	routes.ItemRoute(app)

	app.Listen(":6001")
}
