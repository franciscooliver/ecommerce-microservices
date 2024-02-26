package rabbitmq

import (
	"context"

	ampq "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*ampq.Channel, error) {
	conn, err := ampq.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *ampq.Channel, out chan ampq.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-payment",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ctx context.Context, ch *ampq.Channel, body, exchangeName string) error {
	err := ch.PublishWithContext(
		ctx,
		exchangeName,
		"PaymentDone",
		false,
		false,
		ampq.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
