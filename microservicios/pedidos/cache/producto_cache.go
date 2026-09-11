package cache

import (
	"sync"
	"time"

	"pedidos/models"
)

// ProductoCache es una caché simple en memoria con expiración (TTL)
// para evitar golpear el repositorio en cada consulta de GET /productos.
type ProductoCache struct {
	mu        sync.RWMutex
	productos []models.Producto
	expiresAt time.Time
	ttl       time.Duration
}

func NewProductoCache(ttl time.Duration) *ProductoCache {
	return &ProductoCache{ttl: ttl}
}

func (c *ProductoCache) Obtener() ([]models.Producto, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.productos == nil || time.Now().After(c.expiresAt) {
		return nil, false
	}
	return c.productos, true
}

func (c *ProductoCache) Guardar(productos []models.Producto) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.productos = productos
	c.expiresAt = time.Now().Add(c.ttl)
}
