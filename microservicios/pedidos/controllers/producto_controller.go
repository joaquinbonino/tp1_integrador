package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pedidos/services"
)

// ProductoController traduce las peticiones HTTP en llamadas al servicio de productos.
type ProductoController struct {
	service *services.ProductoService
}

func NewProductoController(service *services.ProductoService) *ProductoController {
	return &ProductoController{service: service}
}

func (pc *ProductoController) ListarProductos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"productos": pc.service.ListarProductos()})
}
