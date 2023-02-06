package order

import "sp-interview/internal/fiber/models"

type AddOrderDto struct {
	OrderID int    `json:"order_id" validate:"required"`
	Price   int    `json:"price" validate:"required"`
	Title   string `json:"title" validate:"required"`
}

func ConvertToModel(o_dto *AddOrderDto) *models.Order {
	return &models.Order{
		OrderID: o_dto.OrderID,
		Price:   o_dto.Price,
		Title:   o_dto.Title,
	}
}
