package redis

import (
	"context"
	"encoding/json"
	"strconv"

	"products-info-service/internal/models"
	"products-info-service/internal/pkg/cache"
)

type QuantityRepo struct {
	redis *cache.Redis
}

func NewQuantityRepo(redis *cache.Redis) *QuantityRepo {
	return &QuantityRepo{redis: redis}
}

func (q *QuantityRepo) AddOrUpdateQuantity(ctx context.Context, quantity *models.Quantity) (bool, error) {
	err := q.redis.Client.Set(ctx, strconv.Itoa(quantity.ID), quantity, 0).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (q *QuantityRepo) GetProductQuantity(ctx context.Context, id int) (*models.Quantity, error) {
	val, err := q.redis.Client.Get(ctx, strconv.Itoa(id)).Result()
	if err != nil {
		return nil, err
	}
	var quantity models.Quantity
	err = json.Unmarshal([]byte(val), &quantity)
	if err != nil {
		return nil, err
	}
	return &quantity, nil
}

func (q *QuantityRepo) DeleteProduct(ctx context.Context, id int) error {
	_, err := q.redis.Client.Del(ctx, strconv.Itoa(id)).Result()
	return err
}
