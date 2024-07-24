package services

import (
	"order-ms/consumer/domain"
	"order-ms/consumer/repositories"
)

type OrderService struct {
	repository *repositories.OrderRepository
}

func (service *OrderService) Create(email, status string, total float32) (*domain.Orders, error) {
	order, err := service.repository.CreateOrder(email, status, total)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderService() *OrderService {
	repository := repositories.NewOrderRepository()
	return &OrderService{
		repository: repository,
	}
}
