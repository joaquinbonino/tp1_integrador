package messaging

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"pedidos/models"
)

// QueueName es la cola donde pedidos publica el evento pedido.confirmado.
// Logística (fuera del alcance de este TP) sería quien la consuma.
const QueueName = "pedidos-confirmados"

// RabbitMQPublisher encapsula la conexión y el canal usados para publicar eventos.
type RabbitMQPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQPublisher(amqpURI string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	if _, err := ch.QueueDeclare(QueueName, true, false, false, false, nil); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQPublisher{conn: conn, ch: ch}, nil
}

func (p *RabbitMQPublisher) PublicarPedidoConfirmado(evento models.PedidoConfirmado) error {
	body, err := json.Marshal(evento)
	if err != nil {
		return err
	}

	err = p.ch.PublishWithContext(context.Background(), "", QueueName, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
	if err != nil {
		return err
	}

	log.Printf("Evento publicado en %s: %s", QueueName, body)
	return nil
}

func (p *RabbitMQPublisher) Close() {
	p.ch.Close()
	p.conn.Close()
}
