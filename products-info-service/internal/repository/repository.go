//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import (
	"context"
	"errors"

	"products-info-service/internal/models"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

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
