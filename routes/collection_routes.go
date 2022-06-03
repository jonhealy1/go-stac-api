package routes

import (
	"go-stac-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/collections", controllers.CreateCollection)
	app.Get("/collections/:collectionId", controllers.GetACollection)
	app.Put("/collections/:collectionId", controllers.EditACollection)
	app.Delete("/collections/:collectionId", controllers.DeleteACollection)
	app.Get("/collections", controllers.GetAllCollections)
}
