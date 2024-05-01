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

/*

type ProductsRepo interface {
	Add(ctx context.Context, product *models.Product) (int, error)
	GetById(ctx context.Context, id int) (*models.Product, error)
	List(ctx context.Context) ([]*models.ProductLightweight, error)
	Update(ctx context.Context, user *models.Product) (bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}

type QuantityRepo interface {
	AddOrUpdateQuantity(ctx context.Context, quantity *models.Quantity) (bool, error)
	GetProductQuantity(ctx context.Context, id int) (*models.Quantity, error)
	DeleteProduct(ctx context.Context, id int) error
}

*/

func (pis *ProductInfoService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	// Logic to get a product by ID from PostgresSQL or Redis
	product, err := pis.ProductsRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pis *ProductInfoService) AddProduct(ctx context.Context, product models.Product) (int, error) {
	return 0, nil
}
