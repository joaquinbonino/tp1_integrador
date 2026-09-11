package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"pedidos/cache"
	"pedidos/controllers"
	"pedidos/messaging"
	"pedidos/repositories"
	"pedidos/services"
)

// amqpURI apunta al RabbitMQ levantado con el docker-compose de CLASE_4
// (usuario/contraseña user/pass en localhost:5672).
const amqpURI = "amqp://user:pass@localhost:5672"

func main() {
	publisher, err := messaging.NewRabbitMQPublisher(amqpURI)
	if err != nil {
		log.Fatalf("no se pudo conectar a RabbitMQ (¿está levantado con docker compose?): %v", err)
	}
	defer publisher.Close()

	productoRepo := repositories.NewProductoMemoryRepository()
	productoCache := cache.NewProductoCache(30 * time.Second)
	productoService := services.NewProductoService(productoRepo, productoCache)
	productoController := controllers.NewProductoController(productoService)

	pedidoRepo := repositories.NewPedidoMemoryRepository()
	pedidoService := services.NewPedidoService(pedidoRepo, productoService, publisher)
	pedidoController := controllers.NewPedidoController(pedidoService)

	router := gin.Default()
	router.GET("/productos", productoController.ListarProductos)
	router.POST("/pedidos", pedidoController.ConfirmarPedido)

	log.Println("Microservicio de pedidos escuchando en http://localhost:8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
