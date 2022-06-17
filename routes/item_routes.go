package routes

import (
	"go-stac-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func ItemRoute(app *fiber.App) {
	app.Post("/collections/:collectionId/items", controllers.CreateItem)
	// app.Get("/items/:collectionId", controllers.GetItem)
	// app.Put("/items/:collectionId", controllers.EditItem)
	// app.Delete("/items/:collectionId", controllers.DeleteItem)
	// app.Get("/items", controllers.GetItems)
}
