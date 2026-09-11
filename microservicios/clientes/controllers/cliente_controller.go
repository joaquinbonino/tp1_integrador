package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"clientes/repositories"
	"clientes/services"
)

// ClienteController traduce las peticiones HTTP en llamadas al servicio de clientes.
type ClienteController struct {
	service *services.ClienteService
}

func NewClienteController(service *services.ClienteService) *ClienteController {
	return &ClienteController{service: service}
}

type crearClienteRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Email  string `json:"email"`
}

func (cc *ClienteController) CrearCliente(c *gin.Context) {
	var req crearClienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de cliente inválidos"})
		return
	}

	cliente := cc.service.CrearCliente(req.Nombre, req.Email)
	c.JSON(http.StatusCreated, cliente)
}

func (cc *ClienteController) ObtenerCliente(c *gin.Context) {
	id := c.Param("id")

	cliente, err := cc.service.ObtenerCliente(id)
	if err != nil {
		if errors.Is(err, repositories.ErrClienteNoEncontrado) {
			c.JSON(http.StatusNotFound, gin.H{"error": "cliente no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el cliente"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
