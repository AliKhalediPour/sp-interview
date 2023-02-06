package repositories

import (
	"sp-consumer/internal/gorm/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Add(*models.Order) error
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository) Add(order *models.Order) error {
	err := o.DB.Create(order).Error

	if err != nil {
		return err
	}
	return nil
}
