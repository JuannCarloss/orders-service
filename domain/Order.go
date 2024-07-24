package domain

import (
	"github.com/pborman/uuid"
)

type Orders struct {
	ID             string  `json:"id"`
	Customer_email string  `json:"email"`
	Timestamp      string  `json:"date"`
	Status         string  `json:"status"`
	Total          float32 `json:"total"`
}

func NewOrder() *Orders {
	return &Orders{
		ID: uuid.New(),
	}
}
