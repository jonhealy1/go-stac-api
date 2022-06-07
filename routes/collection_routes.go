package routes

import (
	"go-stac-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/collections", controllers.CreateCollection)
	app.Get("/collections/:collectionId", controllers.GetCollection)
	app.Put("/collections/:collectionId", controllers.EditCollection)
	app.Delete("/collections/:collectionId", controllers.DeleteCollection)
	app.Get("/collections", controllers.GetCollections)
}
