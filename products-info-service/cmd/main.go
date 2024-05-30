package main

import (
	"context"
	"log"
	"os"

	"products-info-service/internal/controllers"
	"products-info-service/internal/pkg/cache"
	"products-info-service/internal/pkg/db"
	"products-info-service/internal/repository/postgresql"
	"products-info-service/internal/repository/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ctx := context.Background()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbOps, err := db.NewDB(ctx, dbHost, dbPort, dbName, dbUser, dbPass)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisCache := cache.NewRedis(redisHost, redisPassword)

	productsRepo := postgresql.NewProductsRepo(dbOps)
	quantityRepo := redis.NewQuantityRepo(redisCache)

	controller := controllers.NewProductController(productsRepo, quantityRepo)
	router.GET("/products/:id", controller.GetProduct)
	router.POST("/products", controller.AddProduct)
	router.GET("/list", controller.GetAllProducts)
	router.GET("/quantity/:id", controller.GetProductQuantity)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
