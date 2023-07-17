package suscriptor

import (
	"github.com/k1ngd00m/amz_sync_database/aplicacion/manejador"
	"github.com/k1ngd00m/amz_sync_database/infra/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUE_SYNC_CATALOGO = "sync_catalogo"

type SuscriptorRegistrarProducto struct {
	rabbitMQConsumer           config.RabbitMQConsumer
	manejadorRegistrarProducto manejador.ManejadorRegistrarProducto
}

func NewSuscriptorRegistrarProducto(rabbitMQConsumer *config.RabbitMQConsumer,
	manejadorRegistrarProducto *manejador.ManejadorRegistrarProducto) SuscriptorRegistrarProducto {

	suscriptorRegistrarProducto := *&SuscriptorRegistrarProducto{
		rabbitMQConsumer:           *rabbitMQConsumer,
		manejadorRegistrarProducto: *manejadorRegistrarProducto,
	}

	rabbitMQConsumer.On(QUEUE_SYNC_CATALOGO, suscriptorRegistrarProducto.ejecutar)

	return suscriptorRegistrarProducto
}

func (m *SuscriptorRegistrarProducto) ejecutar(deliveries <-chan amqp.Delivery) {

}
