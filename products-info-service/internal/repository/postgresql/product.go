package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"products-info-service/internal/models"
	"products-info-service/internal/repository"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dbOps interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	GetPool(ctx context.Context) *pgxpool.Pool
}

type ProductsRepo struct {
	db dbOps
}

func NewProductsRepo(db dbOps) *ProductsRepo {
	return &ProductsRepo{db: db}
}

// Add specific product
func (r *ProductsRepo) Add(ctx context.Context, product *models.Product) (int, error) {
	var id int
	err := r.db.ExecQueryRow(ctx,
		`INSERT INTO products(name, description, photo, price, quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		product.Name, product.Description, product.Photo, product.Price, product.Quantity).Scan(&id)
	return id, err
}

// GetById Get product info by id
func (r *ProductsRepo) GetById(ctx context.Context, id int) (*models.Product, error) {
	var product models.Product
	err := r.db.Get(ctx, &product,
		`SELECT id,name,description,photo,price,quantity FROM products WHERE id=$1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrObjectNotFound
	}
	return &product, nil
}

// List of all product light
func (r *ProductsRepo) List(ctx context.Context) ([]*models.ProductLightweight, error) {
	products := make([]*models.ProductLightweight, 0)
	err := r.db.Get(ctx, &products, `SELECT id,name FROM products`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrObjectNotFound
	}
	return products, nil
}

// Update product info
func (r *ProductsRepo) Update(ctx context.Context, product *models.Product) (bool, error) {
	res, err := r.db.Exec(ctx,
		`UPDATE products SET name = $1, description = $2, photo = $3, price = $4, quantity = $5 WHERE id=$6`,
		product.Name, product.Description, product.Photo, product.Price, product.Quantity, product.ID)
	return res.RowsAffected() > 0, err
}

// Delete product by id
func (r *ProductsRepo) Delete(ctx context.Context, id int) (bool, error) {
	res, err := r.db.Exec(ctx, "DELETE FROM products WHERE id=$1", id)
	return res.RowsAffected() > 0, err
}
