package controllers

import (
	"context"
	"net/http"
	"strconv"

	"products-info-service/internal/models"
	"products-info-service/internal/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

type ProductController struct {
	productsRepo repository.ProductsRepo
	quantityRepo repository.QuantityRepo
}

func NewProductController(productsRepo repository.ProductsRepo, quantityRepo repository.QuantityRepo) *ProductController {
	return &ProductController{
		productsRepo: productsRepo,
		quantityRepo: quantityRepo,
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

func (h *ProductController) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type of id, should be int"})
		return
	}
	ctx := context.Background()
	product, err := h.productsRepo.GetById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductController) AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	id, err := h.productsRepo.Add(ctx, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	quant := models.Quantity{
		ID:    id,
		Name:  product.Name,
		Count: product.Quantity,
		Price: product.Price,
	}
	_, err = h.quantityRepo.AddOrUpdateQuantity(ctx, &quant)
	if err != nil {
		log.Warningf(ctx, "error", "error while adding to cache")
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *ProductController) GetAllProducts(c *gin.Context) {
	// need implement
	c.JSON(http.StatusOK, _)
}

func (h *ProductController) GetProductQuantity(c *gin.Context) {
	// need implement
	c.JSON(http.StatusOK, _)
}
