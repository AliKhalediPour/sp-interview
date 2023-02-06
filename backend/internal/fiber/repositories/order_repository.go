package repositories

import (
	"context"

	"sp-interview/internal/fiber/models"

	redis_handler "sp-interview/internal/redis"
)

type OrderRepository interface {
	Add(m *models.Order, context context.Context) (*models.Order, error)
}

type orderRepository struct {
	redis redis_handler.RedisHandler
}

func NewOrderRepository(r redis_handler.RedisHandler) OrderRepository {
	return &orderRepository{
		redis: r,
	}
}

func (o *orderRepository) Add(m *models.Order, context context.Context) (*models.Order, error) {
	return nil, nil
}
