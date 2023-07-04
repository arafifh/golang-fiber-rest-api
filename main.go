package main

import (
	"log"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/arafifh/go-rest-api/controller"
	"github.com/arafifh/go-rest-api/repository"
	"github.com/arafifh/go-rest-api/route"
	"github.com/arafifh/go-rest-api/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create database instance
	db := client.Database("golang-fiber") // Replace "your-db-name" with your actual database name

	// Initialize repositories, services, and controllers
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	// Setup routes
	route.SetupProductRoutes(app, productController)

	// Start the server
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}