package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"products-info-service/internal/models"
	"products-info-service/internal/repository"

	"github.com/gin-gonic/gin"
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
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type of id, should be int"})
		return
	}
	product, err := h.productsRepo.GetById(c, id)
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
	id, err := h.productsRepo.Add(c, &product)
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
	_, err = h.quantityRepo.AddOrUpdateQuantity(c, &quant)
	if err != nil {
		log.Print("error while adding to redis: ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *ProductController) GetAllProducts(c *gin.Context) {
	allProducts, err := h.productsRepo.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(allProducts)
	c.JSON(http.StatusOK, allProducts)

}

func (h *ProductController) GetProductQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type of id, should be int"})
		return
	}
	quantity, err := h.quantityRepo.GetProductQuantity(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quantity)
}
