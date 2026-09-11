package services

import (
	"pedidos/cache"
	"pedidos/models"
	"pedidos/repositories"
)

// ProductoService concentra las reglas de negocio del catálogo de productos,
// incluyendo el uso de la caché para el listado.
type ProductoService struct {
	repo  repositories.ProductoRepository
	cache *cache.ProductoCache
}

func NewProductoService(repo repositories.ProductoRepository, cache *cache.ProductoCache) *ProductoService {
	return &ProductoService{repo: repo, cache: cache}
}

func (s *ProductoService) ListarProductos() []models.Producto {
	if productos, ok := s.cache.Obtener(); ok {
		return productos
	}

	productos := s.repo.Listar()
	s.cache.Guardar(productos)
	return productos
}

func (s *ProductoService) BuscarPorID(id string) (models.Producto, bool) {
	return s.repo.BuscarPorID(id)
}
