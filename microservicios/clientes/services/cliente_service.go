package services

import (
	"clientes/models"
	"clientes/repositories"
)

// ClienteService concentra las reglas de negocio del microservicio de clientes.
type ClienteService struct {
	repo repositories.ClienteRepository
}

func NewClienteService(repo repositories.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) CrearCliente(nombre, email string) models.Cliente {
	cliente := models.Cliente{Nombre: nombre, Email: email}
	return s.repo.Guardar(cliente)
}

func (s *ClienteService) ObtenerCliente(id string) (models.Cliente, error) {
	return s.repo.BuscarPorID(id)
}
