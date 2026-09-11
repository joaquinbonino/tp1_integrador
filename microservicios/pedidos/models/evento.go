package models

// PedidoConfirmado es el evento que se publica en la cola pedidos-confirmados
// cuando el servicio de pedidos confirma una compra. Logística es el
// destinatario conceptual de este evento.
type PedidoConfirmado struct {
	Tipo       string `json:"tipo"`
	PedidoID   string `json:"pedido_id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
}
