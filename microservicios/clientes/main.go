package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"clientes/controllers"
	"clientes/repositories"
	"clientes/services"
)

func main() {
	repo := repositories.NewClienteMemoryRepository()
	service := services.NewClienteService(repo)
	controller := controllers.NewClienteController(service)

	router := gin.Default()
	router.POST("/clientes", controller.CrearCliente)
	router.GET("/clientes/:id", controller.ObtenerCliente)

	log.Println("Microservicio de clientes escuchando en http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
