package handlers

import (
	"net/http"
	"strconv"

	"products-info-service/internal/models"
	"products-info-service/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.ProductInfoService
}

func NewHandler(service *services.ProductInfoService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type of id, should be int"})
		return
	}
	product, err := h.service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
