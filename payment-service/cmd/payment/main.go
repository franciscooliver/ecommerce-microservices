package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/franciscooliver/microservice/payment/internal/entity"
	"github.com/franciscooliver/microservice/payment/pkg/rabbitmq"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx := context.Background()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan ampq.Delivery)
	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		var orderRequest entity.OrderRequest
		err := json.Unmarshal(msg.Body, &orderRequest)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		response, err := orderRequest.Process()
		if err != nil {
			slog.Error(err.Error())
			break
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		err = rabbitmq.Publish(ctx, ch, string(responseJSON), "amq.direct")
		if err != nil {
			slog.Error(err.Error())
			break
		}
		msg.Ack(false)
		slog.Info("Ordem Processada")
	}
}
