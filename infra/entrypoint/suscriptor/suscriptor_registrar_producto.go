package suscriptor

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/k1ngd00m/amz_sync_database/aplicacion/dto"
	"github.com/k1ngd00m/amz_sync_database/aplicacion/manejador"
	"github.com/k1ngd00m/amz_sync_database/infra/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUE_SYNC_CATALOGO = "sync_catalogo"

const HEADER_ACCION_ACTUALIZAR_PRODUCTO = "actualizar_producto"
const HEADER_ACCION_REGISTRAR_PRODUCTO = "registrar_producto"

type SuscriptorRegistrarProducto struct {
	rabbitMQConsumer            config.RabbitMQConsumer
	manejadorRegistrarProducto  manejador.ManejadorRegistrarProducto
	manejadorActualizarProducto manejador.ManejadorActualizarProducto
}

func NewSuscriptorRegistrarProducto(rabbitMQConsumer *config.RabbitMQConsumer,
	manejadorRegistrarProducto *manejador.ManejadorRegistrarProducto, manejadorActualizarProducto *manejador.ManejadorActualizarProducto) SuscriptorRegistrarProducto {

	suscriptorRegistrarProducto := *&SuscriptorRegistrarProducto{
		rabbitMQConsumer:            *rabbitMQConsumer,
		manejadorRegistrarProducto:  *manejadorRegistrarProducto,
		manejadorActualizarProducto: *manejadorActualizarProducto,
	}

	rabbitMQConsumer.On(QUEUE_SYNC_CATALOGO, suscriptorRegistrarProducto.ejecutar)

	return suscriptorRegistrarProducto
}

func (m *SuscriptorRegistrarProducto) ejecutar(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {

		str1 := bytes.NewBuffer(delivery.Body).String()
		fmt.Println("String =", str1)

		accion := delivery.Headers["accion"]

		var body dto.DtoProducto
		err := json.Unmarshal(delivery.Body, &body)

		if err != nil {
			fmt.Println(err)
		}

		if accion == HEADER_ACCION_REGISTRAR_PRODUCTO {

			err = m.manejadorRegistrarProducto.Ejecutar(&body)

		} else if accion == HEADER_ACCION_ACTUALIZAR_PRODUCTO {

			err = m.manejadorActualizarProducto.Ejecutar(&body)
		}

		if err != nil {
			fmt.Println(err.Error())
			delivery.Ack(false)
			return
		}

		delivery.Ack(true)

	}

}
