package services

import (
	"order-ms/consumer/domain"
	"order-ms/consumer/handlers"
	"order-ms/consumer/repositories"
)

type OrderService struct {
	repository *repositories.OrderRepository
}

func (service *OrderService) Create(email, status, date string, total float32) (*domain.Orders, error) {
	order, err := service.repository.CreateOrder(email, status, date, total)
	handlers.CheckErr(err)

	return order, nil
}

func NewOrderService() *OrderService {
	repository := repositories.NewOrderRepository()
	return &OrderService{
		repository: repository,
	}
}
