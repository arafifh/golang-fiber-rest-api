package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arafifh/go-rest-api/controller"
)

func SetupProductRoutes(app *fiber.App, productController *controller.ProductController) {
	products := app.Group("/products")

	products.Get("/", productController.GetProducts)
	products.Get("/:id", productController.GetProductByID)
	products.Post("/", productController.CreateProduct)
	products.Put("/:id", productController.UpdateProduct)
	products.Delete("/:id", productController.DeleteProduct)
}
