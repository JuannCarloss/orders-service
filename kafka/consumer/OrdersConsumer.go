package consumer

import (
	"encoding/json"
	"log"
	"order-ms/consumer/domain"
	"order-ms/consumer/handlers"
	"order-ms/consumer/services"

	"github.com/IBM/sarama"
)

var (
	broker = []string{"localhost:9092"}
	topic  = "order"
)

func ConsumeOrders() {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(broker, config)
	handlers.CheckErr(err)

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	handlers.CheckErr(err)

	channel := make(chan struct{})

	svc := services.NewOrderService()

	go func() {
		for {

			select {
			case err := <-consumer.Errors():
				handlers.CheckErr(err)

			case msg := <-consumer.Messages():

				order := domain.NewOrder()
				err := json.Unmarshal([]byte(msg.Value), &order)

				handlers.CheckErr(err)

				svc.Create(order.Customer_email, order.Status, order.Total)

				log.Println("Received messages", string(msg.Value))
			}
		}
	}()
	<-channel
}
