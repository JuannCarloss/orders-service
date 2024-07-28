package repositories

import (
	"order-ms/consumer/config"
	"order-ms/consumer/domain"
	"order-ms/consumer/handlers"

	"github.com/gocql/gocql"
)

type OrderRepository struct {
	db *gocql.Session
}

func (repository *OrderRepository) CreateOrder(email, status, timestamp string, total float32) (*domain.Orders, error) {
	order := domain.NewOrder()
	order.Customer_email = email
	order.Timestamp = timestamp
	order.Status = status
	order.Total = total

	if err := repository.db.Query(`
        INSERT INTO orders (email, date, status, total_price)
        VALUES (?, ?, ?, ?)`,
		order.Customer_email, order.Timestamp, order.Status, order.Total).Exec(); err != nil {

		handlers.CheckErr(err)
	}

	return order, nil
}

func NewOrderRepository() *OrderRepository {
	db, _ := config.InitCluster()

	return &OrderRepository{
		db: db,
	}
}
