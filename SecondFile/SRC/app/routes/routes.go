package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/you/fiber-crud/SRC/app/handlers"
)

func ItemRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/items", handlers.CreateItem)
	api.Get("/items", handlers.GetAllItems)
	api.Get("/items/:id", handlers.GetItemById)
	api.Put("/items/:id", handlers.Update_items_By_ID)
	api.Delete("/items/:id", handlers.DeleteItems)

}
