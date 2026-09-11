package repositories

import "pedidos/models"

// ProductoRepository define el acceso a datos de productos.
type ProductoRepository interface {
	Listar() []models.Producto
	BuscarPorID(id string) (models.Producto, bool)
}

// ProductoMemoryRepository guarda el catálogo de productos en memoria.
type ProductoMemoryRepository struct {
	productos []models.Producto
}

func NewProductoMemoryRepository() *ProductoMemoryRepository {
	return &ProductoMemoryRepository{
		productos: []models.Producto{
			{ID: "P-1", Nombre: "Auriculares", Stock: 10},
			{ID: "P-2", Nombre: "Teclado", Stock: 8},
		},
	}
}

func (r *ProductoMemoryRepository) Listar() []models.Producto {
	return r.productos
}

func (r *ProductoMemoryRepository) BuscarPorID(id string) (models.Producto, bool) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Producto{}, false
}
