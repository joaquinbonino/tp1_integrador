package repositories

import (
	"fmt"
	"sync"

	"pedidos/models"
)

// PedidoRepository define el acceso a datos de pedidos.
type PedidoRepository interface {
	Guardar(pedido models.Pedido) models.Pedido
}

// PedidoMemoryRepository guarda los pedidos confirmados en memoria.
type PedidoMemoryRepository struct {
	mu       sync.Mutex
	pedidos  map[string]models.Pedido
	contador int
}

func NewPedidoMemoryRepository() *PedidoMemoryRepository {
	return &PedidoMemoryRepository{pedidos: make(map[string]models.Pedido)}
}

func (r *PedidoMemoryRepository) Guardar(pedido models.Pedido) models.Pedido {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contador++
	pedido.ID = fmt.Sprintf("PED-%d", r.contador)
	pedido.Estado = "confirmado"
	r.pedidos[pedido.ID] = pedido
	return pedido
}
