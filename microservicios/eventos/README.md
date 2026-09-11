# Eventos de e-commerce

Esta carpeta documenta el contrato de eventos del sistema. El evento mínimo
del trabajo es `pedido.confirmado`, publicado por el microservicio `pedidos`
en la cola `pedidos-confirmados` cuando confirma una compra.

El destinatario conceptual es logística, que en un caso real prepararía el
envío. No se implementa consumidor ni microservicio de logística en este TP.

La struct que define el evento vive junto a quien lo publica, en
[`microservicios/pedidos/models/evento.go`](../pedidos/models/evento.go),
para mantener a cada microservicio dueño de su propio contrato:

```json
{
  "tipo": "pedido.confirmado",
  "pedido_id": "PED-1",
  "cliente_id": "C-1",
  "producto_id": "P-1"
}
```
