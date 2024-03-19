package main

import (
	"context"
	"log"
	"strconv"

	"products-info-service/internal/handlers"
	"products-info-service/internal/pkg/cache"
	"products-info-service/internal/pkg/db"
	"products-info-service/internal/repository/postgresql"
	"products-info-service/internal/repository/redis"
	"products-info-service/internal/services"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ctx := context.Background()
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort, _ := strconv.Atoi(viper.GetString("database.port"))
	dbName := viper.GetString("database.name")
	dbOps, err := db.NewDB(ctx, dbHost, dbPort, dbName, dbUser, dbPassword)
	if err != nil {
		log.Fatal("oh(")
	}
	productsRepo := postgresql.NewProductsRepo(dbOps)
	redisHost := viper.GetString("redis.host")
	redisPassword := viper.GetString("redis.password")
	redisCache := cache.NewRedis(redisHost, redisPassword)
	quantityRepo := redis.NewQuantityRepo(redisCache)
	service := services.NewProductInfoService(productsRepo, quantityRepo)
	handler := handlers.NewHandler(service)
	router.GET("/products/:id", handler.GetProduct)
	router.POST("/products", handler.AddProduct)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
