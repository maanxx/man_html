package routes

import (
	"web-fiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	productGroup := app.Group("/products")

	productGroup.Get("/", handlers.GetProducts)

	productGroup.Post("/", handlers.CreateProduct)

	productGroup.Get("/:id", handlers.GetProductById)

	productGroup.Delete("/:id", handlers.DeleteProduct)
}