package repositories

import (
	"fmt"
	"order-ms/consumer/config"
	"order-ms/consumer/domain"
	"time"

	"github.com/gocql/gocql"
)

type OrderRepository struct {
	db *gocql.Session
}

func (repository *OrderRepository) CreateOrder(email, status string, total float32) (*domain.Orders, error) {
	order := domain.NewOrder()
	order.Customer_email = email
	order.Timestamp = time.Now().Format("2006-01-02T15:04:05")
	order.Status = status
	order.Total = total

	if err := repository.db.Query(`
        INSERT INTO orders (id, email, date, status, total_price)
        VALUES (?, ?, ?, ?, ?)`,
		order.ID, order.Customer_email, order.Timestamp, order.Status, order.Total).Exec(); err != nil {

		fmt.Println("Error while inserting data: ", err)
		return nil, err
	}

	return order, nil
}

func NewOrderRepository() *OrderRepository {
	db, _ := config.InitCluster()

	return &OrderRepository{
		db: db,
	}
}
