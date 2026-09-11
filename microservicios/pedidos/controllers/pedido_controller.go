package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"pedidos/services"
)

// PedidoController traduce las peticiones HTTP en llamadas al servicio de pedidos.
type PedidoController struct {
	service *services.PedidoService
}

func NewPedidoController(service *services.PedidoService) *PedidoController {
	return &PedidoController{service: service}
}

type confirmarPedidoRequest struct {
	ClienteID  string `json:"cliente_id" binding:"required"`
	ProductoID string `json:"producto_id" binding:"required"`
}

func (pc *PedidoController) ConfirmarPedido(c *gin.Context) {
	var req confirmarPedidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de pedido inválidos"})
		return
	}

	pedido, err := pc.service.ConfirmarPedido(req.ClienteID, req.ProductoID)
	if err != nil {
		if errors.Is(err, services.ErrProductoInexistente) {
			c.JSON(http.StatusNotFound, gin.H{"error": "producto inexistente"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo confirmar el pedido"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "pedido confirmado",
		"pedido":  pedido,
	})
}
