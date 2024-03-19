package services

import (
	"context"

	"products-info-service/internal/models"
	"products-info-service/internal/repository"
)

type ProductInfoService struct {
	ProductsRepo repository.ProductsRepo
	QuantityRepo repository.QuantityRepo
}

func NewProductInfoService(ProductsRepo repository.ProductsRepo, QuantityRepo repository.QuantityRepo) *ProductInfoService {
	return &ProductInfoService{
		ProductsRepo: ProductsRepo,
		QuantityRepo: QuantityRepo,
	}
}

func (pis *ProductInfoService) GetProductByID(id int) (*models.Product, error) {
	// Logic to get a product by ID from PostgresSQL or Redis
	// Example PostgresSQL query execution
	ctx := context.Background()
	product, err := pis.ProductsRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pis *ProductInfoService) AddProduct(product models.Product) (int, error) {
	return 0, nil
}
