package services

import (
	"errors"
	"log"

	"pedidos/messaging"
	"pedidos/models"
	"pedidos/repositories"
)

// ErrProductoInexistente se devuelve al intentar confirmar un pedido de un producto que no existe.
var ErrProductoInexistente = errors.New("producto inexistente")

// PedidoService concentra las reglas de negocio de la confirmación de pedidos.
type PedidoService struct {
	pedidoRepo      repositories.PedidoRepository
	productoService *ProductoService
	publisher       *messaging.RabbitMQPublisher
}

func NewPedidoService(pedidoRepo repositories.PedidoRepository, productoService *ProductoService, publisher *messaging.RabbitMQPublisher) *PedidoService {
	return &PedidoService{pedidoRepo: pedidoRepo, productoService: productoService, publisher: publisher}
}

func (s *PedidoService) ConfirmarPedido(clienteID, productoID string) (models.Pedido, error) {
	if _, ok := s.productoService.BuscarPorID(productoID); !ok {
		return models.Pedido{}, ErrProductoInexistente
	}

	pedido := models.Pedido{ClienteID: clienteID, ProductoID: productoID}
	pedido = s.pedidoRepo.Guardar(pedido)

	evento := models.PedidoConfirmado{
		Tipo:       "pedido.confirmado",
		PedidoID:   pedido.ID,
		ClienteID:  pedido.ClienteID,
		ProductoID: pedido.ProductoID,
	}

	if err := s.publisher.PublicarPedidoConfirmado(evento); err != nil {
		log.Printf("no se pudo publicar el evento pedido.confirmado: %v", err)
	}

	return pedido, nil
}
