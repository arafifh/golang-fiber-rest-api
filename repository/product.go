package repository

import (
	"context"
	"log"

	"github.com/arafifh/go-rest-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(database *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: database.Collection("products"),
	}
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	ctx := context.TODO()
	cur, err := r.collection.Find(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)

	var products []model.Product
	err = cur.All(ctx, &products)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id string) (*model.Product, error) {
	ctx := context.TODO()
	var product model.Product

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(product model.Product) (*model.Product, error) {
	ctx := context.TODO()
	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(id string, product model.Product) (*model.Product, error) {
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": objID}, product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id string) error {
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
