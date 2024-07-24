package main

import (
	"order-ms/consumer/services"
)

func main() {

	svc := services.NewOrderService()

	svc.Create("teste", "AWAITING_PAYMENT", 100)
}
