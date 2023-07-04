package service

import (
	"log"

	"github.com/arafifh/go-rest-api/model"
	"github.com/arafifh/go-rest-api/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id string) (*model.Product, error) {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func (s *ProductService) CreateProduct(product model.Product) (*model.Product, error) {
	createdProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return createdProduct, nil
}

func (s *ProductService) UpdateProduct(id string, product model.Product) (*model.Product, error) {
	updatedProduct, err := s.productRepository.UpdateProduct(id, product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	err := s.productRepository.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}