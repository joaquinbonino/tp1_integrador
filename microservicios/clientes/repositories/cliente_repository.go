package repositories

import (
	"errors"
	"fmt"
	"sync"

	"clientes/models"
)

// ErrClienteNoEncontrado se devuelve cuando no existe un cliente con el id buscado.
var ErrClienteNoEncontrado = errors.New("cliente no encontrado")

// ClienteRepository define el acceso a datos de clientes, independiente del almacenamiento usado.
type ClienteRepository interface {
	Guardar(cliente models.Cliente) models.Cliente
	BuscarPorID(id string) (models.Cliente, error)
}

// ClienteMemoryRepository guarda los clientes en un mapa en memoria.
type ClienteMemoryRepository struct {
	mu       sync.RWMutex
	clientes map[string]models.Cliente
	contador int
}

func NewClienteMemoryRepository() *ClienteMemoryRepository {
	return &ClienteMemoryRepository{
		clientes: make(map[string]models.Cliente),
	}
}

func (r *ClienteMemoryRepository) Guardar(cliente models.Cliente) models.Cliente {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contador++
	cliente.ID = fmt.Sprintf("C-%d", r.contador)
	r.clientes[cliente.ID] = cliente
	return cliente
}

func (r *ClienteMemoryRepository) BuscarPorID(id string) (models.Cliente, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cliente, ok := r.clientes[id]
	if !ok {
		return models.Cliente{}, ErrClienteNoEncontrado
	}
	return cliente, nil
}
