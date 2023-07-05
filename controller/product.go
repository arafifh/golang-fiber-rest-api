package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/arafifh/go-rest-api/model"
	"github.com/arafifh/go-rest-api/service"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) GetProducts(ctx *fiber.Ctx) error {
	products, err := c.productService.GetProducts()
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	
	}
	return ctx.JSON(products)
}

func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := c.productService.GetProductByID(id)
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get product",
		})
	}

	return ctx.JSON(product)
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	product := new(model.Product)
	
	if err := ctx.BodyParser(product); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}

	createdProduct, err := c.productService.CreateProduct(*product)
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return ctx.JSON(createdProduct)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product := new(model.Product)
	
	if err := ctx.BodyParser(product); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}

	updatedProduct, err := c.productService.UpdateProduct(id, *product)
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return ctx.JSON(updatedProduct)
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.productService.DeleteProduct(id)
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}